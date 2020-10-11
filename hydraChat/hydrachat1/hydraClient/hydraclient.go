package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	done := make(chan string)
	connection, _ := net.Dial("tcp", "localhost:2300")
	go listenToServer(connection)

	<-done

}

func listenToServer(connection net.Conn) {
	scanner := bufio.NewScanner(connection)

	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println("received message: ", text)
	}

}

func writeToServer(connection net.Conn) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Fprintf(connection, text)
	}
}
