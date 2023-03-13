package main

import (
	"fmt"
	"net"
	"os"
	"testing"
	"time"
)

func listenForConn(server net.Listener) {
	incConn, err := server.Accept()
	if err != nil {
		fmt.Println("Error accepting: ", err.Error())
		os.Exit(1)
	}
	// Handle connections in a new goroutine.
	handleRequest(incConn)
}
func TestTcpConnection(t *testing.T) {
	connString := "localhost:3333"
	server, _ := net.Listen("tcp", connString)
	defer server.Close()

	go listenForConn(server)
	outConn, _ := connectToSocket(connString)
	time.Sleep(2 * time.Second)
	testWrite := []byte("WOW WOW WOW\r\n")
	outConn.Write(testWrite)
	time.Sleep(2 * time.Second)
	outConn.Close()
}

func handleRequest(conn net.Conn) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	reqLen, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	fmt.Println(reqLen)
	fmt.Println(buf)
	// Send a response back to person contacting us.
	conn.Write([]byte("Message received."))
	// Close the connection when you're done with it.
	conn.Close()
}
