package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func receiveUserInput(readerPtr io.Reader) (string, error) {
	fmt.Print("Enter text: ")
	reader := bufio.NewReader(readerPtr)
	// ReadString will block until the delimiter is entered
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return "", errors.New("Could not read input")
	}
	// remove the delimeter from the string
	input = strings.TrimSuffix(input, "\n")
	return input, nil
}

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
	userInput, _ := receiveUserInput(os.Stdin)
	fmt.Println("User inputed:" + userInput)
}
