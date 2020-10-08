package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	done := make(chan string)
	connection, _ := net.Dial("tcp", ":2300")

	go listenForMessages(connection)
	go writeMessages(connection)
	<-done
}

func listenForMessages(connection net.Conn) {
	scanner := bufio.NewScanner(connection)

	for scanner.Scan() {
		msg := scanner.Text()
		fmt.Println("message Received: ", msg)
	}

}

func writeMessages(connection net.Conn) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		msg := scanner.Text()
		fmt.Fprintf(connection, msg+"\n")
	}
}
