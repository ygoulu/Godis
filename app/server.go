package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

type event struct {
	conn    net.Conn
	command []byte
}

var eventQueue = make(chan event)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage

	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	defer l.Close()

	go func() {
		for {
			connection, err := l.Accept()
			if err != nil {
				fmt.Println("Error accepting connection: ", err.Error())
				continue
			}

			go handleConn(connection)
		}
	}()

	//recieve executions of the event loop
	eventLoop()
}

func eventLoop() {
	for {
		e := <-eventQueue
		e.conn.Write([]byte("+PONG\r\n"))
		e.conn.Close()
	}
}

func handleConn(conn net.Conn) {
	log.Println("New connection from", conn.RemoteAddr())

	// Read the request from the connection
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Println("Error reading from connection:", err)
		return
	}

	eventQueue <- event{conn, buf[:n]}
}
