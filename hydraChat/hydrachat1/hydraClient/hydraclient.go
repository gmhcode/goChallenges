package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	done := make(chan string)
	connection, _ := net.Dial("tcp", "127.0.0.1:2300")
	writeToServer(connection)
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
	go func() {
		for scanner.Scan() {
			text := scanner.Text()
			//need the \n
			fmt.Fprintf(connection, text+"\n")
		}
	}()

}
