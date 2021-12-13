package concensus

import (
	"github.com/hashicorp/raft"
	"github.com/rs/zerolog"
	pb "simple-twitter.com/backend/rpc/proto"
)

// Node is a representation of a raft participant
type Node struct {
	Config   *Config
	RaftNode *raft.Raft
	FSM      *FSM
	Log      *zerolog.Logger
}

// Event represents the concensus interaction type and data
type Event struct {
	Type             string
	NewUserRequest   *pb.User
	NewTweetRequest  *pb.Tweet
	NewFollowRequest *pb.Follow
}
