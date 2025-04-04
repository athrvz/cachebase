package main

import (
	"cachebase/internal/config"
	"cachebase/internal/pkg/persistence"
	"cachebase/internal/proto/gen"
	"cachebase/internal/server"
	"cachebase/internal/utilities"
	"fmt"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	// Load .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	// Create a new gRPC server
	grpcServer := grpc.NewServer()
	fmt.Println(os.Getenv("CONFIG_FILE_PATH"))
	// Load rdb config
	config := config.LoadConfig(os.Getenv("CONFIG_FILE_PATH"))
	// if _, err := os.Stat(config.RDB.FilePath); err != nil {
	// 	fmt.Println("err with file")
	// }

	// Load previously stored RDB
	if err := persistence.LoadRDB(config.RDB); err != nil {
		log.Printf("RDB load failed: %v, Starting with empty dataset.", err)
	} else {
		log.Printf("Loaded %d keys from RDB Snapshot", utilities.GetCachedKeysCount())
	}

	stopRDB := make(chan struct{})
	persistence.StartRDBSaver(config.RDB, stopRDB) // Background Saver
	defer close(stopRDB)                           // Clean shutdown
	// Handle shutdown signals
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	//

	//

	// Register the CacheServer with the gRPC server
	cacheServer := server.NewCacheServer()
	gen.RegisterCacheServiceServer(grpcServer, cacheServer)

	// Start listening on a port
	lis, err := net.Listen("tcp", ":"+os.Getenv("GRPC_PORT"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("gRPC server is running on port 50051...")

	// Start the gRPC server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
