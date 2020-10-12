package main

import (
	"bufio"
	"net"
)

func main() {
	done := make(chan bool)
	listener, _ := net.Listen("tcp", ":2300")

	go startListening(listener)

	<-done
}

func startListening(listener net.Listener) {
	for {
		conn, _ := listener.Accept()
		writer := bufio.NewWriter(conn)
		// _ = writer
		go func() {

			scanner := bufio.NewScanner(conn)

			for scanner.Scan() {
				msg := scanner.Text()
				writer.WriteString(msg)
				writer.Flush()

			}
		}()

	}
}

func sendMessages(writer bufio.Writer) {

}
