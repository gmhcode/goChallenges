package simple

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

//ClientStart - Starts the Clients tcp connection
func ClientStart() net.Conn {
	conn, err := net.Dial("tcp", "127.0.0.1:2300")
	if err != nil {
		log.Fatal("Could not connect to hydra chat system", err)
	}
	clientScanMessagesFromServer(conn)
	clientBeginSending(conn)
	return conn
}

func clientScanMessagesFromServer(conn net.Conn) {
	scanner := bufio.NewScanner(conn)

	go func() {
		for scanner.Scan() {
			fmt.Println("Message Received: ", scanner.Text())

		}
	}()
}

func clientBeginSending(connection net.Conn) {
	scanner := bufio.NewScanner(os.Stdin)
	go func() {
		for scanner.Scan() {
			fmt.Fprintf(connection, scanner.Text())

		}

	}()

}

func ServerStartListenAndAccept(connection string) {

	listener, _ := net.Listen("tcp", connection)

	go func() {
		for {
			conn, _ := listener.Accept()
			hanldeConnection(conn)
		}
	}()

}

func hconnectionanldeConnection(connection net.Conn) {
	fmt.Println("Received Connection From: ", connection.RemoteAddr())

}
