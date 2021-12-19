package rpc

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	"simple-twitter.com/backend/rpc/proto"
	"testing"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	rpc.RegisterTwitterServer(s, &twitterServer{})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func Test_CreateUser(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := rpc.NewTwitterClient(conn)
	resp, err := client.CreateUser(ctx, &rpc.User{Fullname: "Alex", Password: "alex123", Handle: "@alex", Bio: "Hey there!"})
	if err != nil {
		t.Fatalf("CreateUser failed: %v", err)
	}
	log.Printf("Response: %+v", resp)
	// Test for output here.
	var expected = rpc.UserReply{User: &rpc.User{Id: "1", Fullname: "Alex", Password: "alex123", Handle: "@alex", Bio: "Hey there!"} }
	if exp, got := &expected, resp; exp != got {
		t.Fatalf("unexpected results for query, expected %s, got %s", exp, got)
	}
}

func Test_LoginUser(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := rpc.NewTwitterClient(conn)
	resp, err := client.LoginUser(ctx, &rpc.User{Fullname: "Alex", Password: "alex123"})
	if err != nil {
		t.Fatalf("CreateUser failed: %v", err)
	}
	log.Printf("Response: %+v", resp)
	// Test for output here.
	var expected = rpc.UserReply{User: &rpc.User{Id: "1", Fullname: "Alex", Password: "alex123", Handle: "@alex", Bio: "Hey there!"} }
	if exp, got := &expected, resp; exp != got {
		t.Fatalf("unexpected results for query, expected %s, got %s", exp, got)
	}
}

func Test_CreateTweet(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := rpc.NewTwitterClient(conn)
	resp, err := client.CreateTweet(ctx, &rpc.Tweet{Content: "Lets see what how this Raft holds up...", AuthorId: "3", AuthorName: "Alex", AuthorHandle: "@alext"})
	if err != nil {
		t.Fatalf("CreateUser failed: %v", err)
	}
	log.Printf("Response: %+v", resp)
	// Test for output here.
	var tweets = []*rpc.Tweet{{Id: "1", Content: "Lets see what how this Raft holds up...", AuthorId: "3", AuthorName: "Alex", AuthorHandle: "@alext"}}
	var expected = rpc.TweetsReply{Tweet: tweets }
	if exp, got := &expected, resp; exp != got {
		t.Fatalf("unexpected results for query, expected %s, got %s", exp, got)
	}
}
