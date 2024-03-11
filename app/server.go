package main

import (
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	log.Println("Server listening on port 8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	log.Println("New connection from", conn.RemoteAddr())

	// Read the request from the connection
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Println("Error reading from connection:", err)
		return
	}
	request := string(buf[:n])

	// Check if the request is a ping
	if request == "ping" {
		// Respond with a pong
		_, err = conn.Write([]byte("pong"))
		if err != nil {
			log.Println("Error writing to connection:", err)
			return
		}
		log.Println("Sent pong response")
	} else {
		// Respond with an error for unrecognized requests
		_, err = conn.Write([]byte("Error: Unrecognized request"))
		if err != nil {
			log.Println("Error writing to connection:", err)
			return
		}
		log.Println("Sent error response for unrecognized request")
	}
}
