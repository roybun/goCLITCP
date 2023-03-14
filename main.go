package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
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

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, os.Interrupt)
	go func() {
		sig := <-sigs
		fmt.Println("Recieved signal:")
		fmt.Println(sig)
		done <- true
		os.Exit(0)
		return
	}()

	for {
		userInput, _ := receiveUserInput(os.Stdin)
		fi, _ := os.Create("test.txt")
		fmt.Println("User inputed:" + userInput)
		fi.Write([]byte(userInput))
		fi.Close()
		//Check if signal was recieved
		select {
		case <-done:
			fmt.Println("exiting")
			return
		default:
		}
	}

}
