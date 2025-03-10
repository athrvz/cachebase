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
	value := req.GetValue() // value is byte slice
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
