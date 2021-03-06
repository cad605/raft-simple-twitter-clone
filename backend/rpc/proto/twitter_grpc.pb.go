// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package rpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TwitterClient is the client API for Twitter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TwitterClient interface {
	// Handles join requests for a new node
	HandleJoin(ctx context.Context, in *JoinRaftRequest, opts ...grpc.CallOption) (*JoinRaftReply, error)
	// Creates a new user
	CreateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*UserReply, error)
	// Returns authentication for a given username and password
	LoginUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*UserReply, error)
	// Creates a new tweet authored by a given user with the given content
	CreateTweet(ctx context.Context, in *Tweet, opts ...grpc.CallOption) (*TweetsReply, error)
	// Allows a user to add another user to their list of followers
	FollowUser(ctx context.Context, in *Follow, opts ...grpc.CallOption) (*FollowReply, error)
	// Allows a user to remove a user from their list of followers
	UnfollowUser(ctx context.Context, in *Follow, opts ...grpc.CallOption) (*FollowReply, error)
	// Returns user info for a given user
	GetUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*UserReply, error)
	// Returns a list of users
	GetUsers(ctx context.Context, in *User, opts ...grpc.CallOption) (*ManyUsersReply, error)
	// Returns tweets authored by a given user
	GetTweetsByUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*TweetsReply, error)
	// Returns tweets authored by those users followed by a given user
	GetFeedByUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*TweetsReply, error)
	// Returns a list of people that follow a given user
	GetFollowedByUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*ManyUsersReply, error)
	// Returns the list of users that a given user follows
	GetFollowingByUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*ManyUsersReply, error)
	GetUsersNotFollowed(ctx context.Context, in *User, opts ...grpc.CallOption) (*ManyUsersReply, error)
}

type twitterClient struct {
	cc grpc.ClientConnInterface
}

func NewTwitterClient(cc grpc.ClientConnInterface) TwitterClient {
	return &twitterClient{cc}
}

func (c *twitterClient) HandleJoin(ctx context.Context, in *JoinRaftRequest, opts ...grpc.CallOption) (*JoinRaftReply, error) {
	out := new(JoinRaftReply)
	err := c.cc.Invoke(ctx, "/Twitter/HandleJoin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *twitterClient) CreateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*UserReply, error) {
	out := new(UserReply)
	err := c.cc.Invoke(ctx, "/Twitter/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *twitterClient) LoginUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*UserReply, error) {
	out := new(UserReply)
	err := c.cc.Invoke(ctx, "/Twitter/LoginUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *twitterClient) CreateTweet(ctx context.Context, in *Tweet, opts ...grpc.CallOption) (*TweetsReply, error) {
	out := new(TweetsReply)
	err := c.cc.Invoke(ctx, "/Twitter/CreateTweet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *twitterClient) FollowUser(ctx context.Context, in *Follow, opts ...grpc.CallOption) (*FollowReply, error) {
	out := new(FollowReply)
	err := c.cc.Invoke(ctx, "/Twitter/FollowUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *twitterClient) UnfollowUser(ctx context.Context, in *Follow, opts ...grpc.CallOption) (*FollowReply, error) {
	out := new(FollowReply)
	err := c.cc.Invoke(ctx, "/Twitter/UnfollowUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *twitterClient) GetUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*UserReply, error) {
	out := new(UserReply)
	err := c.cc.Invoke(ctx, "/Twitter/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *twitterClient) GetUsers(ctx context.Context, in *User, opts ...grpc.CallOption) (*ManyUsersReply, error) {
	out := new(ManyUsersReply)
	err := c.cc.Invoke(ctx, "/Twitter/GetUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *twitterClient) GetTweetsByUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*TweetsReply, error) {
	out := new(TweetsReply)
	err := c.cc.Invoke(ctx, "/Twitter/GetTweetsByUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *twitterClient) GetFeedByUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*TweetsReply, error) {
	out := new(TweetsReply)
	err := c.cc.Invoke(ctx, "/Twitter/GetFeedByUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *twitterClient) GetFollowedByUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*ManyUsersReply, error) {
	out := new(ManyUsersReply)
	err := c.cc.Invoke(ctx, "/Twitter/GetFollowedByUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *twitterClient) GetFollowingByUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*ManyUsersReply, error) {
	out := new(ManyUsersReply)
	err := c.cc.Invoke(ctx, "/Twitter/GetFollowingByUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *twitterClient) GetUsersNotFollowed(ctx context.Context, in *User, opts ...grpc.CallOption) (*ManyUsersReply, error) {
	out := new(ManyUsersReply)
	err := c.cc.Invoke(ctx, "/Twitter/GetUsersNotFollowed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TwitterServer is the server API for Twitter service.
// All implementations must embed UnimplementedTwitterServer
// for forward compatibility
type TwitterServer interface {
	// Handles join requests for a new node
	HandleJoin(context.Context, *JoinRaftRequest) (*JoinRaftReply, error)
	// Creates a new user
	CreateUser(context.Context, *User) (*UserReply, error)
	// Returns authentication for a given username and password
	LoginUser(context.Context, *User) (*UserReply, error)
	// Creates a new tweet authored by a given user with the given content
	CreateTweet(context.Context, *Tweet) (*TweetsReply, error)
	// Allows a user to add another user to their list of followers
	FollowUser(context.Context, *Follow) (*FollowReply, error)
	// Allows a user to remove a user from their list of followers
	UnfollowUser(context.Context, *Follow) (*FollowReply, error)
	// Returns user info for a given user
	GetUser(context.Context, *User) (*UserReply, error)
	// Returns a list of users
	GetUsers(context.Context, *User) (*ManyUsersReply, error)
	// Returns tweets authored by a given user
	GetTweetsByUser(context.Context, *User) (*TweetsReply, error)
	// Returns tweets authored by those users followed by a given user
	GetFeedByUser(context.Context, *User) (*TweetsReply, error)
	// Returns a list of people that follow a given user
	GetFollowedByUser(context.Context, *User) (*ManyUsersReply, error)
	// Returns the list of users that a given user follows
	GetFollowingByUser(context.Context, *User) (*ManyUsersReply, error)
	GetUsersNotFollowed(context.Context, *User) (*ManyUsersReply, error)
	mustEmbedUnimplementedTwitterServer()
}

// UnimplementedTwitterServer must be embedded to have forward compatible implementations.
type UnimplementedTwitterServer struct {
}

func (UnimplementedTwitterServer) HandleJoin(context.Context, *JoinRaftRequest) (*JoinRaftReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleJoin not implemented")
}
func (UnimplementedTwitterServer) CreateUser(context.Context, *User) (*UserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedTwitterServer) LoginUser(context.Context, *User) (*UserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}
func (UnimplementedTwitterServer) CreateTweet(context.Context, *Tweet) (*TweetsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTweet not implemented")
}
func (UnimplementedTwitterServer) FollowUser(context.Context, *Follow) (*FollowReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FollowUser not implemented")
}
func (UnimplementedTwitterServer) UnfollowUser(context.Context, *Follow) (*FollowReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnfollowUser not implemented")
}
func (UnimplementedTwitterServer) GetUser(context.Context, *User) (*UserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedTwitterServer) GetUsers(context.Context, *User) (*ManyUsersReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}
func (UnimplementedTwitterServer) GetTweetsByUser(context.Context, *User) (*TweetsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTweetsByUser not implemented")
}
func (UnimplementedTwitterServer) GetFeedByUser(context.Context, *User) (*TweetsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFeedByUser not implemented")
}
func (UnimplementedTwitterServer) GetFollowedByUser(context.Context, *User) (*ManyUsersReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFollowedByUser not implemented")
}
func (UnimplementedTwitterServer) GetFollowingByUser(context.Context, *User) (*ManyUsersReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFollowingByUser not implemented")
}
func (UnimplementedTwitterServer) GetUsersNotFollowed(context.Context, *User) (*ManyUsersReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsersNotFollowed not implemented")
}
func (UnimplementedTwitterServer) mustEmbedUnimplementedTwitterServer() {}

// UnsafeTwitterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TwitterServer will
// result in compilation errors.
type UnsafeTwitterServer interface {
	mustEmbedUnimplementedTwitterServer()
}

func RegisterTwitterServer(s grpc.ServiceRegistrar, srv TwitterServer) {
	s.RegisterService(&Twitter_ServiceDesc, srv)
}

func _Twitter_HandleJoin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinRaftRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TwitterServer).HandleJoin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Twitter/HandleJoin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TwitterServer).HandleJoin(ctx, req.(*JoinRaftRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Twitter_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TwitterServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Twitter/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TwitterServer).CreateUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _Twitter_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TwitterServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Twitter/LoginUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TwitterServer).LoginUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _Twitter_CreateTweet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Tweet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TwitterServer).CreateTweet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Twitter/CreateTweet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TwitterServer).CreateTweet(ctx, req.(*Tweet))
	}
	return interceptor(ctx, in, info, handler)
}

func _Twitter_FollowUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Follow)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TwitterServer).FollowUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Twitter/FollowUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TwitterServer).FollowUser(ctx, req.(*Follow))
	}
	return interceptor(ctx, in, info, handler)
}

func _Twitter_UnfollowUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Follow)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TwitterServer).UnfollowUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Twitter/UnfollowUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TwitterServer).UnfollowUser(ctx, req.(*Follow))
	}
	return interceptor(ctx, in, info, handler)
}

func _Twitter_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TwitterServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Twitter/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TwitterServer).GetUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _Twitter_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TwitterServer).GetUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Twitter/GetUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TwitterServer).GetUsers(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _Twitter_GetTweetsByUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TwitterServer).GetTweetsByUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Twitter/GetTweetsByUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TwitterServer).GetTweetsByUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _Twitter_GetFeedByUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TwitterServer).GetFeedByUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Twitter/GetFeedByUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TwitterServer).GetFeedByUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _Twitter_GetFollowedByUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TwitterServer).GetFollowedByUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Twitter/GetFollowedByUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TwitterServer).GetFollowedByUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _Twitter_GetFollowingByUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TwitterServer).GetFollowingByUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Twitter/GetFollowingByUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TwitterServer).GetFollowingByUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _Twitter_GetUsersNotFollowed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TwitterServer).GetUsersNotFollowed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Twitter/GetUsersNotFollowed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TwitterServer).GetUsersNotFollowed(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

// Twitter_ServiceDesc is the grpc.ServiceDesc for Twitter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Twitter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Twitter",
	HandlerType: (*TwitterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HandleJoin",
			Handler:    _Twitter_HandleJoin_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _Twitter_CreateUser_Handler,
		},
		{
			MethodName: "LoginUser",
			Handler:    _Twitter_LoginUser_Handler,
		},
		{
			MethodName: "CreateTweet",
			Handler:    _Twitter_CreateTweet_Handler,
		},
		{
			MethodName: "FollowUser",
			Handler:    _Twitter_FollowUser_Handler,
		},
		{
			MethodName: "UnfollowUser",
			Handler:    _Twitter_UnfollowUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _Twitter_GetUser_Handler,
		},
		{
			MethodName: "GetUsers",
			Handler:    _Twitter_GetUsers_Handler,
		},
		{
			MethodName: "GetTweetsByUser",
			Handler:    _Twitter_GetTweetsByUser_Handler,
		},
		{
			MethodName: "GetFeedByUser",
			Handler:    _Twitter_GetFeedByUser_Handler,
		},
		{
			MethodName: "GetFollowedByUser",
			Handler:    _Twitter_GetFollowedByUser_Handler,
		},
		{
			MethodName: "GetFollowingByUser",
			Handler:    _Twitter_GetFollowingByUser_Handler,
		},
		{
			MethodName: "GetUsersNotFollowed",
			Handler:    _Twitter_GetUsersNotFollowed_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/twitter.proto",
}
