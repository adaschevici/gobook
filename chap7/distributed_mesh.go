package main

import (
	"fmt"
	"net"
)

type Subscriber struct {
	Addres     net.Addr
	Connection net.Conn
	do         chan Task
}

type Task struct {
	name string
}

var SubscriberCount int
var Subscribers []Subscriber
var CurrentSubscriber int
var taskChannel chan Task

func (sb Subscriber) awaitTask() {
	select {
	case t := <-sb.do:
		fmt.Println("Performing task:", t.name)
	}
}
func broadcast() {
	for i := range Subscribers {
		for j := range Subscribers {
			Subscribers[i].Connection.Write
			([]byte("Subscriber:", Subscribers[j].Address))
		}
	}
}

func serverListen(listener net.Listener) {
	for {
		conn, _ := listener.Accept()

		SubscriberCount++

		subscriber := Subscriber{Address: conn.remoteAddr(), Connection: conn}

		subscriber.do = make(chan Task)

		subscriber.awaitTask()

		_ = append(Subscribers, subscriber)

		broadcast()
	}
}

func doTask() {
	for {
		select {
		case task := <-taskChannel:
			fmt.Println(task.Name, "invoked")
			Subscribers[CurrentSubscriber].do <- task
			if (CurrentSubscriber + 1) > SubscriberCount {
				CurrentSubscriber = 0
			} else {
				CurrentSubscriber++
			}
		}
	}
}

func main() {

	destinationStatus := make(chan int)

	SubscriberCount = 0
	CurrentSubscriber = 0

	taskChannel = make(chan Task)

	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		fmt.Println("Could not start server!", err)
	}
	go serverListen(listener)
	go doTask()

	<-destinationStatus
}
