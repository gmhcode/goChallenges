package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	done := make(chan bool)
	listener, _ := net.Listen("tcp", ":2300")

	startListening(listener)

	<-done
}

func startListening(listener net.Listener) {
	for {
		conn, _ := listener.Accept()
		writer := bufio.NewWriter(conn)
		fmt.Println("New connection made")
		// _ = writer
		go func() {

			newReader := bufio.NewReader(conn)
			scanner := bufio.NewScanner(newReader)
			fmt.Println("hit")
			for scanner.Scan() {
				msg := scanner.Text()
				fmt.Println("received msg : ", msg)
				writer.WriteString(msg)
				writer.Flush()

			}
		}()

	}
}

func sendMessages(writer bufio.Writer) {

}
