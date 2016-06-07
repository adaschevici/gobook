package main

import (
	"fmt"
	"time"
)

func thinkForAWhile() {
	for {
		fmt.Println("Still thinking")
		time.Sleep(1 * time.Second)
	}
}

func main() {
	fmt.Println("Where did I leave my keys?")

	blockChannel := make(chan int)
	go thinkForAWhile()

	<-blockChannel

	fmt.Println("Foun'em!!!")
}
