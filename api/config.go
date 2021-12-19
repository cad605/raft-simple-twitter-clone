package main

import (
	"fmt"
	multierror "github.com/hashicorp/go-multierror"
	template "github.com/hashicorp/go-sockaddr/template"
	flag "github.com/ogier/pflag"
	"net"
)

type rawConfig struct {
	RaftAddress string
	RaftPort    int
	ServerPort    int
}

type config struct {
	ServerAddress net.Addr
	RaftAddress net.Addr
}

type configError struct {
	ConfigurationPoint string
	Err                error
}

func (err *configError) Error() string {
	return fmt.Sprintf("%s: %s", err.ConfigurationPoint, err.Err.Error())
}

func resolveRawConfig(rawConfig *rawConfig) (*config, error) {
	var errors *multierror.Error

	// Bind address
	var raftAddress net.IP
	resolvedBindAddr, err := template.Parse(rawConfig.RaftAddress)
	if err != nil {
		configErr := &configError{
			ConfigurationPoint: "raft-address",
			Err:                err,
		}
		errors = multierror.Append(errors, configErr)
	} else {
		raftAddress = net.ParseIP(resolvedBindAddr)
		if raftAddress == nil {
			err := fmt.Errorf("cannot parse IP address: %s", resolvedBindAddr)
			configErr := &configError{
				ConfigurationPoint: "raft-address",
				Err:                err,
			}
			errors = multierror.Append(errors, configErr)
		}
	}

	// Raft port
	if rawConfig.RaftPort < 1 || rawConfig.RaftPort > 65536 {
		configErr := &configError{
			ConfigurationPoint: "raft-port",
			Err:                fmt.Errorf("port numbers must be 1 < port < 65536"),
		}
		errors = multierror.Append(errors, configErr)
	}

	// Construct Raft Address
	raftAddr := &net.TCPAddr{
		IP:   raftAddress,
		Port: rawConfig.RaftPort,
	}

	// HTTP port
	if rawConfig.ServerPort < 1 || rawConfig.ServerPort > 65536 {
		configErr := &configError{
			ConfigurationPoint: "http-port",
			Err:                fmt.Errorf("port numbers must be 1 < port < 65536"),
		}
		errors = multierror.Append(errors, configErr)
	}

	// Construct HTTP Address
	serverAddr := &net.TCPAddr{
		Port: rawConfig.ServerPort,
	}

	if err := errors.ErrorOrNil(); err != nil {
		return nil, err
	}

	return &config{
		RaftAddress: raftAddr,
		ServerAddress: serverAddr,
	}, nil
}

func getRawConfig() *rawConfig {
	var config rawConfig

	flag.StringVarP(&config.RaftAddress, "rafter-address", "r",
		"127.0.0.1", "IP Address on which to bind")

	flag.IntVarP(&config.RaftPort, "raft-port", "b",
		8000, "Port on which to bind Raft")

	flag.IntVarP(&config.ServerPort, "server-port", "p",
		8080, "Port on which to bind Server")

	flag.Parse()
	return &config
}
