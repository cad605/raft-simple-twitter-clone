package main

import (
	"context"
	"fmt"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/raft"
	"github.com/hashicorp/raft-boltdb"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"path/filepath"
	"simple-twitter.com/backend/concensus"
	"simple-twitter.com/backend/rpc"
	pb "simple-twitter.com/backend/rpc/proto"
	"time"
)

func main() {
	logger := zerolog.New(os.Stdout)
	ctx := context.Background()

	// Read in a raft node configuration
	rawConfig := concensus.ReadRawConfig()
	config, err := concensus.ResolveConfig(rawConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Configuration errors - %s\n", err)
		os.Exit(1)
	}

	// Create a new node
	nodeLogger := logger.With().Str("component", "Node").Logger()
	node, err := NewNode(config, &nodeLogger)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error configuring Node: %s", err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stderr, "Setting up rpc on - %s\n", config.HTTPAddress.String())
	lis, err := net.Listen("tcp", config.HTTPAddress.String())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to listen at - %s: %s\n", config.HTTPAddress.String(), err)
		os.Exit(1)
	}

	if config.JoinAddress != "" {
		var opts []grpc.DialOption
		opts = append(opts, grpc.WithInsecure())
		conn, err := grpc.Dial(config.JoinAddress, opts...)
		if err != nil {
			log.Fatalf("fail to dial: %v", err)
		}
		defer conn.Close()
		client := pb.NewTwitterClient(conn)
		if _, err := client.HandleJoin(ctx, &pb.JoinRaftRequest{PeerAddress: config.RaftAddress.String()}); err != nil {
			fmt.Fprintf(os.Stderr, "Unable to join at - %s: %s\n", config.RaftAddress.String(), err)
			return
		}
	}

	grpcServer := grpc.NewServer()
	s := rpc.NewTwitterServer(node, &nodeLogger)
	pb.RegisterTwitterServer(grpcServer, s)

	err = grpcServer.Serve(lis)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to start grpcServer at - %s\n", err)
		os.Exit(1)
		return
	}

	fmt.Fprintf(os.Stderr, "Started Raft successfully...\n")
	terminate := make(chan os.Signal, 1)
	grpcServer.GracefulStop()
	lis.Close()
	signal.Notify(terminate, os.Interrupt)
	<-terminate
	fmt.Fprintf(os.Stderr, "Exiting Raft successfully...\n")

}

func NewNode(config *concensus.Config, log *zerolog.Logger) (*concensus.Node, error) {

	if err := os.MkdirAll(filepath.Join(config.DataDir, config.RaftAddress.String()), 0700); err != nil {
		return nil, err
	}

	// Create raft configuration.
	raftConfig := raft.DefaultConfig()
	raftConfig.LocalID = raft.ServerID(config.RaftAddress.String())
	raftConfig.Logger = hclog.Default()

	// Create raft transport.
	transportLogger := log.With().Str("component", "raft-transport").Logger()
	transport, err := raftTransport(config.RaftAddress, transportLogger)
	if err != nil {
		return nil, err
	}

	// Create the snapshot store. This allows Raft to truncate the log.
	snapshotStoreLogger := log.With().Str("component", "raft-snapshots").Logger()
	snapshotStore, err := raft.NewFileSnapshotStore(config.DataDir, 1, snapshotStoreLogger)
	if err != nil {
		return nil, err
	}

	// Create the log store and stable store.
	logStore, err := raftboltdb.NewBoltStore(filepath.Join(config.DataDir, config.RaftAddress.String(), "raft-log.bolt"))
	if err != nil {
		return nil, err
	}
	stableStore, err := raftboltdb.NewBoltStore(filepath.Join(config.DataDir, config.RaftAddress.String(), "raft-stable.bolt"))
	if err != nil {
		return nil, err
	}

	// We need an in-memory backend, at least for bootstrapping purposes.
	log.Printf("creating in-memory backend at NewNode")
	db, err := concensus.CreateOnDisk(filepath.Join(config.DataDir, config.RaftAddress.String(), "backend.db"), nil)
	if err != nil {
		return nil, err
	}

	// Create backend tables
	concensus.CreateTables(db.RwDB)

	// Create finite state machine
	fsm := &concensus.FSM{
		DB: db,
	}
	log.Printf("created database at: %s", filepath.Join(config.DataDir, config.RaftAddress.String(), "backend.db"))

	// Create a new raft
	raftNode, err := raft.NewRaft(raftConfig, fsm, logStore, stableStore,
		snapshotStore, transport)
	if err != nil {
		return nil, err
	}

	// Bootstrap if set
	if config.Bootstrap {
		configuration := raft.Configuration{
			Servers: []raft.Server{
				{
					ID:      raftConfig.LocalID,
					Address: transport.LocalAddr(),
				},
			},
		}
		raftNode.BootstrapCluster(configuration)
		if err != nil {
			return nil, fmt.Errorf("raft.Raft.BootstrapCluster: %v", err)
		}
	}

	// return node
	return &concensus.Node{
		Config:   config,
		RaftNode: raftNode,
		Log:      log,
		FSM:      fsm,
	}, nil
}

// raftTransport creates the network transport (TCP) for communication between the nodes
func raftTransport(raftAddr net.Addr, log io.Writer) (*raft.NetworkTransport, error) {
	address, err := net.ResolveTCPAddr("tcp", raftAddr.String())
	if err != nil {
		return nil, err
	}

	transport, err := raft.NewTCPTransport(address.String(), address, 3, 10*time.Second, log)
	if err != nil {
		return nil, err
	}

	return transport, nil
}
