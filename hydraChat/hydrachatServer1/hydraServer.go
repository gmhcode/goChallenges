package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	done := make(chan bool)
	listener, _ := net.Listen("tcp", ":2300")
	serverListen(listener)

	<-done
}

func serverListen(listener net.Listener) {
	for {
		conn, _ := listener.Accept()
		fmt.Println("accepted connection")
		go handleConnection(conn)
	}
}

func handleConnection(connection net.Conn) {

	reader := bufio.NewReader(connection)
	writer := bufio.NewWriter(connection)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		msg := scanner.Text()
		fmt.Println(msg)
		distributeReceivedMessages(msg, writer)

	}
}

func scanForMessages(connection net.Conn) {

}

func distributeReceivedMessages(msg string, writer *bufio.Writer) {
	writer.WriteString("weve received your message" + "\n")
	writer.Flush()
}
