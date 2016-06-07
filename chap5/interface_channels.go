package main

import (
	"fmt"
	"strconv"
)

type Messenger interface {
	Relay() string
}

type Message struct {
	status string
}

func (m Message) Relay() string {
	return m.status
}

func alertMessage(w chan Messenger, i int) {
	m := new(Message)
	m.status = "Done with " + strconv.FormatInt(int64(i), 10)
	v <- m
}

func main() {
	msg := make(chan Messenger)

	for i := 0; i < 10; i++ {
		go alertMessage(msg, i)
	}

	select {
	case message := <-msg:
		fmt.Println(message.Relay())
	}
	<-msg
}
