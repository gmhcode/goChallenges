package main

import (
	"bufio"
	"challenges/hydraChat/simplechat/simple"
	"fmt"
	"log"
	"net"
	"os"
)

type room struct {
	MessageCH chan string
	People    map[chan<- string]struct{}
	Quit      chan struct{}
}

func main() {

	room := room{
		MessageCH: make(chan string),
		Quit:      make(chan struct{}),
	}
	connection := simple.ClientStart()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		msg := scanner.Text()
		_, _ = fmt.Fprintf(connection, msg+"\n")
	}

	<-room.Quit

}

func handleMessages(room *room) {
	go func() {
		for msg := range room.MessageCH {
			fmt.Println("message received: ", msg)
		}
		defer close(room.MessageCH)
	}()
}

func scanForMessages(room *room) {
	var msg string

	go func() {
		for {
			fmt.Scan(&msg)

			room.MessageCH <- msg
			if msg == "QUIT" {
				close(room.Quit)
				return
			}
		}

	}()
}

//ClientStart - Starts the Clients tcp connection
func ClientStart() net.Conn {
	conn, err := net.Dial("tcp", "127.0.0.1:2300")
	if err != nil {
		log.Fatal("Could not connect to hydra chat system", err)
	}
	return conn
}

func ScanMessagesFromServer(conn net.Conn) {
	scanner := bufio.NewScanner(conn)

	go func() {
		for scanner.Scan() {
			fmt.Println("Message Received: ", scanner.Text())

		}
	}()
}

func StartServerListenAndAccept(connection string) {

	listener, _ := net.Listen("tcp", connection)

	go func() {
		for {
			conn, _ := listener.Accept()
			HanldeConnection(conn)
		}
	}()

}

func HanldeConnection(connection net.Conn) {
	fmt.Println("Received Connection From: ", connection.RemoteAddr())

}
