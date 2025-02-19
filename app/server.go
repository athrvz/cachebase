package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
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

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	// Read data
	buff := make([]byte, 1024)
	_, err := conn.Read(buff)

	if err != nil {
		conn.Close()
	}
	fmt.Println("buff: ", string(buff))
	conn.Write([]byte("+PONG\r\n"))
}
