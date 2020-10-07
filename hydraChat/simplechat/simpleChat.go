package main

import (
	"fmt"
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

	scanForMessages(&room)
	handleMessages(&room)
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

// func cancelSignal() {

// }
