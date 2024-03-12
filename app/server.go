package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage

	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	connection, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
	handleConn(connection)
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	log.Println("New connection from", conn.RemoteAddr())

	for {
		// Read the request from the connection
		buf := make([]byte, 1024)
		_, err := conn.Read(buf)
		if err != nil {
			log.Println("Error reading from connection:", err)
			return
		}

		// Respond with a pong
		_, err = conn.Write([]byte("+PONG\r\n"))
		if err != nil {
			log.Println("Error writing to connection:", err)
			return
		}
		log.Println("Sent pong response")
	}
}
