package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	connection := startClient()
	done := make(chan struct{})
	//write to server
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			msg := scanner.Text()
			fmt.Fprintf(connection, msg+"\n")
		}
	}()

	//listen for server
	go func() {
		scanner := bufio.NewScanner(connection)

		for scanner.Scan() {
			msg := scanner.Text()
			fmt.Println("Message Received: ", msg)
		}
	}()

	<-done

}

func startClient() net.Conn {
	connection, err := net.Dial("tcp", "localhost:2300")
	if err != nil {
		fmt.Println("error connecting to server", err)
	}
	return connection
}
