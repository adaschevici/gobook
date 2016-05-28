package main

import "fmt"

func doNothing() string {
	return "nothing"
}

func main() {
	var channel chan string = make(chan string)
	go func() {
		a := doNothing()
		channel <- a
	}()
	val := <-channel
	fmt.Println(val)
}
