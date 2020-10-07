package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	quitChan := make(chan struct{})
	connection, _ := net.Dial("tcp", "127.0.0.1:2300")
	startSending(connection)
	go startListening(connection)
	<-quitChan
}

func startSending(connection net.Conn) {
	scanner := bufio.NewScanner(os.Stdin)
	go func() {
		for scanner.Scan() {
			msg := scanner.Text()
			fmt.Fprintf(connection, msg+"\n")
		}
	}()
}

func startListening(connection net.Conn) {
	scanner := bufio.NewScanner(connection)

	for scanner.Scan() {
		msg := scanner.Text()
		fmt.Println("new message received: ", msg)
	}
}
