package main

import (
	"errors"
	"fmt"
	"net"
)

func connectToSocket(address string) (net.Conn, error) {
	//connect to tcp socket
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Could not connect to socket")
	}
	return conn, nil
}

func main() {
	fmt.Println("heelow world")
}
