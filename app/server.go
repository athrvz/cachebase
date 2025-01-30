package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net"
	"os"
)

var _ = net.Listen
var _ = os.Exit

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading the .env file")
	}

	Port := os.Getenv("PORT")

	listener, err := net.Listen("tcp", Port)
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	} else {
		fmt.Println("Listening on port: ", Port)
	}
	defer listener.Close()

	_, err = listener.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

}
