package rpc

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/hashicorp/raft"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"simple-twitter.com/backend/concensus"
	pb "simple-twitter.com/backend/rpc/proto"
	"time"
)

func NewTwitterServer(node *concensus.Node, logger *zerolog.Logger) *twitterServer {
	s := &twitterServer{node: node, logger: logger}
	return s
}

type twitterServer struct {
	pb.UnimplementedTwitterServer
	node   *concensus.Node
	logger *zerolog.Logger
}

func (s *twitterServer) HandleJoin(ctx context.Context, req *pb.JoinRaftRequest) (*pb.JoinRaftReply, error) {
	peerAddress := req.PeerAddress
	if peerAddress == "" {
		s.logger.Error().Msg("Peer-Address not set on request")
		return &pb.JoinRaftReply{Success: false}, errors.New("Peer-Address not set on request")
	}
	s.node.RaftNode.Leader()

	addPeerFuture := s.node.RaftNode.AddVoter(
		raft.ServerID(peerAddress), raft.ServerAddress(peerAddress), 0, 0)
	if err := addPeerFuture.Error(); err != nil {
		s.logger.Error().
			Err(err).
			Str("peer.remoteaddr", peerAddress).
			Msg("Error joining peer to Raft")
		return &pb.JoinRaftReply{Success: false}, err
	}

	s.logger.Info().Str("peer.remoteaddr", peerAddress).Msg("Peer joined Raft")
	return &pb.JoinRaftReply{Success: true}, nil
}

func (s *twitterServer) CreateUser(ctx context.Context, req *pb.User) (*pb.UserReply, error) {

	event := &concensus.Event{
		Type: "CreateUser",
		NewUserRequest: req,
	}

	eventBytes, err := json.Marshal(event)
	if err != nil {
		s.logger.Error().Err(err).Msg("Failed to marshal JSON in CreateUser")
	}

	applyFuture := s.node.RaftNode.Apply(eventBytes, 5*time.Second)
	if err := applyFuture.Error(); err != nil {
		s.logger.Error().Err(err).Msg("Failed to create new user.")
		return &pb.UserReply{}, err
	}

	s.logger.Info().Str("event.createUser", "success").Msg("New user created")
	return &pb.UserReply{}, nil
}

func (s *twitterServer) LoginUser(ctx context.Context, req *pb.User) (*pb.UserReply, error) {
	s.node.FSM.Mutex.Lock()
	defer s.node.FSM.Mutex.Unlock()
	conn, err := s.node.FSM.DB.RoDB.Conn(context.Background())
	if err != nil {
		s.logger.Error().Err(err).Msg("Error opening database connection")
		return &pb.UserReply{}, err
	}
	defer func(conn *sql.Conn) {
		err := conn.Close()
		if err != nil {
			s.logger.Error().Err(err).Msg("Error closing connection.")
			return
		}
	}(conn)

	res, err := s.node.FSM.DB.LoginUser(conn, req)

	if err != nil {
		return &pb.UserReply{}, err
	}

	s.logger.Info().Str("event.login", "success").Msg("Logged in user.")
	return &pb.UserReply{User: res}, nil
}

func (s *twitterServer) CreateTweet(ctx context.Context, req *pb.Tweet) (*pb.TweetsReply, error) {

	event := &concensus.Event{
		Type: "CreateTweet",
		NewTweetRequest: req,
	}

	eventBytes, err := json.Marshal(event)
	if err != nil {
		s.logger.Error().Err(err).Msg("Failed to marshal JSON in CreateTweet")
	}

	applyFuture := s.node.RaftNode.Apply(eventBytes, 5*time.Second)
	if err := applyFuture.Error(); err != nil {
		s.logger.Error().Err(err).Msg("Failed to create new tweet.")
		return &pb.TweetsReply{}, nil
	}

	s.logger.Info().Str("event.createTweet", "success").Msg("New tweet created.")
	return &pb.TweetsReply{}, nil
}

func (s *twitterServer) FollowUser(ctx context.Context, req *pb.Follow) (*pb.FollowReply, error) {

	event := &concensus.Event{
		Type: "FollowUser",
		NewFollowRequest: req,
	}

	eventBytes, err := json.Marshal(event)
	if err != nil {
		s.logger.Error().Err(err).Msg("Failed to marshal JSON in FollowUser")
	}

	applyFuture := s.node.RaftNode.Apply(eventBytes, 5*time.Second)
	if err := applyFuture.Error(); err != nil {
		s.logger.Error().Err(err).Msg("Failed to create new follow.")
		return &pb.FollowReply{Success: false}, nil
	}

	s.logger.Info().Str("event.followUser", "success").Msg("New follow created.")
	return &pb.FollowReply{Success: true}, nil
}

func (s *twitterServer) UnfollowUser(ctx context.Context, req *pb.Follow) (*pb.FollowReply, error) {

	event := &concensus.Event{
		Type: "UnfollowUser",
		NewFollowRequest: req,
	}

	eventBytes, err := json.Marshal(event)
	if err != nil {
		s.logger.Error().Err(err).Msg("Failed to marshal JSON in UnfollowUser")
	}

	applyFuture := s.node.RaftNode.Apply(eventBytes, 5*time.Second)
	if err := applyFuture.Error(); err != nil {
		s.logger.Error().Err(err).Msg("Failed to unfollow user.")
		return &pb.FollowReply{Success: false}, nil
	}

	s.logger.Info().Str("event.unfollowUser", "success").Msg("Old follow deleted.")
	return &pb.FollowReply{Success: true}, nil
}

func (s *twitterServer) GetUser(ctx context.Context, req *pb.User) (*pb.UserReply, error) {
	s.node.FSM.Mutex.Lock()
	defer s.node.FSM.Mutex.Unlock()
	conn, err := s.node.FSM.DB.RoDB.Conn(context.Background())
	if err != nil {
		s.logger.Error().Err(err).Msg("Error opening database connection")
		return &pb.UserReply{}, err
	}
	defer func(conn *sql.Conn) {
		err := conn.Close()
		if err != nil {
			s.logger.Error().Err(err).Msg("Error closing connection.")
			return
		}
	}(conn)

	res, err := s.node.FSM.DB.GetUser(conn, req)

	if err != nil {
		s.logger.Error().Err(err).Msg("Error fetching user")
		return &pb.UserReply{}, err
	}

	s.logger.Info().Str("event.getUser", "success").Msg("Fetched user successfully")
	return &pb.UserReply{User: res}, nil
}

func (s *twitterServer) GetUsers(ctx context.Context, req *pb.User) (*pb.ManyUsersReply, error) {
	s.node.FSM.Mutex.Lock()
	defer s.node.FSM.Mutex.Unlock()
	conn, err := s.node.FSM.DB.RoDB.Conn(context.Background())
	if err != nil {
		s.logger.Error().Err(err).Msg("Error opening database connection")
		return &pb.ManyUsersReply{}, err
	}
	defer func(conn *sql.Conn) {
		err := conn.Close()
		if err != nil {
			s.logger.Error().Err(err).Msg("Error closing connection.")
			return
		}
	}(conn)

	res, err := s.node.FSM.DB.GetUsers(conn)

	if err != nil {
		s.logger.Error().Err(err).Msg("Error fetching user")
		return &pb.ManyUsersReply{}, err
	}

	s.logger.Info().Str("event.getUser", "success").Msg("Fetched user successfully")
	return &pb.ManyUsersReply{Users: res}, nil
}

func (s *twitterServer) GetUsersNotFollowed(ctx context.Context, req *pb.User) (*pb.ManyUsersReply, error) {
	s.node.FSM.Mutex.Lock()
	defer s.node.FSM.Mutex.Unlock()
	conn, err := s.node.FSM.DB.RoDB.Conn(context.Background())
	if err != nil {
		s.logger.Error().Err(err).Msg("Error opening database connection")
		return &pb.ManyUsersReply{}, err
	}
	defer func(conn *sql.Conn) {
		err := conn.Close()
		if err != nil {
			s.logger.Error().Err(err).Msg("Error closing connection.")
			return
		}
	}(conn)

	res, err := s.node.FSM.DB.GetUsersNotFollowed(conn, req)

	if err != nil {
		s.logger.Error().Err(err).Msg("Error fetching users not followed.")
		return &pb.ManyUsersReply{}, err
	}

	s.logger.Info().Str("event.getUsersNotFollowed", "success").Msg("Fetched user successfully")
	return &pb.ManyUsersReply{Users: res}, nil
}

func (s *twitterServer) GetTweetsByUser(ctx context.Context, req *pb.User) (*pb.TweetsReply, error) {
	s.node.FSM.Mutex.Lock()
	defer s.node.FSM.Mutex.Unlock()
	conn, err := s.node.FSM.DB.RoDB.Conn(context.Background())
	if err != nil {
		s.logger.Error().Err(err).Msg("Error opening database connection")
		return &pb.TweetsReply{}, err
	}
	defer func(conn *sql.Conn) {
		err := conn.Close()
		if err != nil {
			s.logger.Error().Err(err).Msg("Error closing connection.")
			return
		}
	}(conn)

	res, err := s.node.FSM.DB.GetTweetsByUser(conn, req)

	if err != nil {
		s.logger.Error().Err(err).Msg("Error fetching user tweets")
		return &pb.TweetsReply{}, err
	}

	s.logger.Info().Str("event.GetTweetsByUser", "success").Msg("Fetched user tweets successfully")
	return &pb.TweetsReply{Tweet: res}, nil
}

func (s *twitterServer) GetFeedByUser(ctx context.Context, req *pb.User) (*pb.TweetsReply, error) {
	s.node.FSM.Mutex.Lock()
	defer s.node.FSM.Mutex.Unlock()
	conn, err := s.node.FSM.DB.RoDB.Conn(context.Background())
	if err != nil {
		s.logger.Error().Err(err).Msg("Error opening database connection")
		return &pb.TweetsReply{}, err
	}
	defer func(conn *sql.Conn) {
		err := conn.Close()
		if err != nil {
			s.logger.Error().Err(err).Msg("Error closing connection.")
			return
		}
	}(conn)

	res, err := s.node.FSM.DB.GetFeedByUser(conn, req)
	if err != nil {
		s.logger.Error().Err(err).Msg("Error fetching tweet feed")
		return &pb.TweetsReply{}, err
	}

	s.logger.Info().Str("event.GetFeedByUser", "success").Msg("Fetched user feed successfully")
	return &pb.TweetsReply{Tweet: res}, nil
}

func (s *twitterServer) GetFollowedByUser(ctx context.Context, req *pb.User) (*pb.ManyUsersReply, error) {
	s.node.FSM.Mutex.Lock()
	defer s.node.FSM.Mutex.Unlock()
	conn, err := s.node.FSM.DB.RoDB.Conn(context.Background())
	if err != nil {
		s.logger.Error().Err(err).Msg("Error opening database connection")
		return &pb.ManyUsersReply{}, err
	}
	defer func(conn *sql.Conn) {
		err := conn.Close()
		if err != nil {
			s.logger.Error().Err(err).Msg("Error closing connection.")
			return
		}
	}(conn)

	res, err := s.node.FSM.DB.GetFollowedByUser(conn, req)
	if err != nil {
		s.logger.Error().Err(err).Msg("Error fetching followed by user")
		return &pb.ManyUsersReply{}, err
	}

	s.logger.Info().Str("event.GetFollowedByUser", "success").Msg("Fetched followed by user successfully")
	return &pb.ManyUsersReply{Users: res}, nil
}

func (s *twitterServer) GetFollowingByUser(ctx context.Context, req *pb.User) (*pb.ManyUsersReply, error) {
	s.node.FSM.Mutex.Lock()
	defer s.node.FSM.Mutex.Unlock()
	conn, err := s.node.FSM.DB.RoDB.Conn(context.Background())
	if err != nil {
		s.logger.Error().Err(err).Msg("Error opening database connection")
		return &pb.ManyUsersReply{}, err
	}
	defer func(conn *sql.Conn) {
		err := conn.Close()
		if err != nil {
			s.logger.Error().Err(err).Msg("Error closing connection.")
			return
		}
	}(conn)

	res, err := s.node.FSM.DB.GetFollowingByUser(conn, req)
	if err != nil {
		s.logger.Error().Err(err).Msg("Error fetching following by user")
		return &pb.ManyUsersReply{}, err
	}

	s.logger.Info().Str("event.GetFollowedByUser", "success").Msg("Fetched following by user successfully")
	return &pb.ManyUsersReply{Users: res}, nil
}
