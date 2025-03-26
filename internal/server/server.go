package server

import (
	"cachebase/internal/commands"
	"cachebase/internal/pkg/storage"
	"cachebase/internal/proto/gen"
	"context"

	// "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type CacheServer struct {
	gen.UnimplementedCacheServiceServer // Embed
}

// New instance of cache server
func NewCacheServer() *CacheServer {
	return &CacheServer{}
}

// Get handles the Get RPC request
func (s *CacheServer) Get(ctx context.Context, req *gen.GetRequest) (*gen.GetResponse, error) {
	key := req.GetKey()
	value, found := commands.Get(key)
	if !found {
		return nil, status.Errorf(codes.NotFound, "key not found: %s", key)
	}
	// Converting value to bytes (assuming value is already a byte slice)
	byteValue, ok := value.(string)
	if !ok {
		return nil, status.Errorf(codes.Internal, "invalid value type")
	}

	return &gen.GetResponse{
		Value: byteValue,
		Found: true,
	}, nil
}

// Ping handles the Ping RPC request
func (s *CacheServer) Ping(ctx context.Context, req *emptypb.Empty) (*gen.PingResponse, error) {
	return &gen.PingResponse{
		Pong: "PONG",
	}, nil
}

// Set handles the Set RPC request
func (s *CacheServer) Set(ctx context.Context, req *gen.SetRequest) (*gen.SetResponse, error) {
	key := req.GetKey()
	value := req.GetValue()
	ttl := int(req.GetTtl())
	// ToDo: grpc support for ttl
	opts := storage.SetOptions{
		Expiry: &ttl,
	}
	commands.Set(key, value, opts)
	return &gen.SetResponse{
		Success: true,
	}, nil
}

// Delete handles the Delete RPC request.
func (s *CacheServer) Delete(ctx context.Context, req *gen.DeleteRequest) (*gen.DeleteResponse, error) {
	key := req.GetKey()
	commands.Delete(key)
	return &gen.DeleteResponse{
		Success: true,
	}, nil
}

// ListPush handles ListPush RPC request
func (s *CacheServer) ListPush(ctx context.Context, req *gen.ListPushRequest) (*gen.ListPushResponse, error) {
	key := req.GetKey()
	values := req.GetValues()
	length := commands.ListPush(key, values)
	return &gen.ListPushResponse{
		Length: int32(length),
	}, nil
}

// ListPop
func (s *CacheServer) ListPop(ctx context.Context, req *gen.ListPopRequest) (*gen.ListPopResponse, error) {
	key := req.GetKey()
	value, exists := commands.ListPop(key)
	if !exists {
		return nil, status.Errorf(codes.NotFound, "key not found of list is empty: %s", key)
	}
	return &gen.ListPopResponse{
		Value: value,
	}, nil
}

// SetAdd function
func (s *CacheServer) SetAdd(ctx context.Context, req *gen.SetAddRequest) (*gen.SetAddResponse, error) {
	key := req.GetKey()
	values := req.GetValues()
	ttl := int(req.GetTtl())
	// ToDo: grpc support for ttl
	opts := storage.SetOptions{
		Expiry: &ttl,
	}
	added := commands.SetAdd(key, values, opts)
	return &gen.SetAddResponse{
		Added: int32(added),
	}, nil
}

func (s *CacheServer) SetRemove(ctx context.Context, req *gen.SetRemoveRequest) (*gen.SetRemoveResponse, error) {
	key := req.GetKey()
	values := req.GetValues()
	removed := commands.SetRemove(key, values)
	return &gen.SetRemoveResponse{
		Removed: int32(removed),
	}, nil
}

// SetIsMember handles the SetIsMember RPC request.
func (s *CacheServer) SetIsMember(ctx context.Context, req *gen.SetIsMemberRequest) (*gen.SetIsMemberResponse, error) {
	key := req.GetKey()
	value := req.GetValue()
	isMember := commands.SetIsMember(key, value)
	return &gen.SetIsMemberResponse{IsMember: isMember}, nil
}

// SetMembers handles the SetMembers RPC request.
func (s *CacheServer) SetMembers(ctx context.Context, req *gen.SetMembersRequest) (*gen.SetMembersResponse, error) {
	key := req.GetKey()
	members, exists := commands.SetMembers(key)
	if !exists {
		return nil, status.Errorf(codes.NotFound, "key not found: %s", key)
	}
	return &gen.SetMembersResponse{Members: members}, nil
}
