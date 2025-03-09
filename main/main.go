package main

import (
	"cachebase/internal/proto/gen"
	"cachebase/internal/server"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	// Register the CacheServer with the gRPC server
	cacheServer := server.NewCacheServer()
	gen.RegisterCacheServiceServer(grpcServer, cacheServer)

	// Start listening on a port
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("gRPC server is running on port 50051...")

	// Start the gRPC server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
