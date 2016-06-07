package main

import (
	"fmt"
	"time"
)

func thinkForAWhile(bC chan int) {
	i := 0
	max := 10
	for {
		fmt.Println(i)
		if i >= max {
			bC <- 1
		}
		fmt.Println("Still thinking")
		time.Sleep(1 * time.Second)
		i++
	}
}

func main() {
	fmt.Println("Where did I leave my keys?")
	blockChannel := make(chan int)
	go thinkForAWhile(blockChannel)

	end := <-blockChannel
	fmt.Println(end)
	fmt.Println("Foun'em!!!")
}
