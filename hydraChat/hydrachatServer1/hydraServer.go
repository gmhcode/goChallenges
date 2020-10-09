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
	writers := make([]*bufio.Writer, 0)
	for {
		conn, _ := listener.Accept()
		newWriter := bufio.NewWriter(conn)
		writers = append(writers, newWriter)
		fmt.Println("accepted connection from")
		go handleConnection(conn, &writers)
	}
}

func handleConnection(connection net.Conn, writers *[]*bufio.Writer) {

	reader := bufio.NewReader(connection)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		msg := scanner.Text()
		fmt.Println(msg)

		distributeReceivedMessages(msg, writers)

	}
}

func distributeReceivedMessages(msg string, writers *[]*bufio.Writer) {
	for _, writer := range *writers {
		writer.WriteString("weve received your message " + msg + "\n")
		writer.Flush()
	}

}
