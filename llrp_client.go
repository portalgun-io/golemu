package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

const (
	CONN_HOST = "0.0.0.0"
	CONN_PORT = "5084"
	CONN_TYPE = "tcp"
	BUFSIZE   = 512
)

func main() {
	// Listen for incoming connections.
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	for {
		// Make a buffer to hold incoming data.
		buf := make([]byte, BUFSIZE)
		// Read the incoming connection into the buffer.
		reqLen, err := conn.Read(buf)
		if err == io.EOF {
			conn.Close()
			break
		} else if err != nil {
			fmt.Println("Error reading:", err.Error())
			fmt.Println("reqLen: " + string(reqLen))
		}
		fmt.Printf("Packet size: %v\n", reqLen)
		fmt.Printf("% x\n", buf[:reqLen])

		// Send a response back to person contacting us.
		//conn.Write([]byte("Message received.\r\n"))
		// Close the connection when you're done with it.
	}
}