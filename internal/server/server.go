package server

import (
	"cachebase/internal/commands"
	"cachebase/internal/proto/gen"
	"context"

	// "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	byteValue, ok := value.([]byte)
	if !ok {
		return nil, status.Errorf(codes.Internal, "invalid value type")
	}

	return &gen.GetResponse{
		Value: byteValue,
		Found: true,
	}, nil
}
