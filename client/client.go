package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// Connect to the server at localhost:6379
	conn, err := net.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}
	defer conn.Close() // Ensure the connection is closed when we're done

	// Send a PING command to the server
	_, err = conn.Write([]byte("PING \"HI\""))
	if err != nil {
		fmt.Println("Error sending PING:", err)
		return
	}

	// Read the response from the server
	response, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// Print the server's response
	fmt.Println("Server response:", response)
}
