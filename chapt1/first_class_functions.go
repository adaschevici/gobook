package main

import (
	"fmt"
	"strings"
)

func shortenString(message string) func() string {
	return func() string {
		messageSlice := strings.Split(message, " ")
		wordLength := len(messageSlice)
		if wordLength < 1 {
			return "Nothing Left!"
		} else {
			messageSlice = messageSlice[:(wordLength - 1)]
			message = strings.Join(messageSlice, " ")
			fmt.Println(wordLength)
			return message
		}
	}
}

func main() {
	myString := shortenString("Welcome to concurrency in Go! ...")

	fmt.Println(myString())
	myString = shortenString("Welcome to concurrency in Go! ...")
	fmt.Println(myString())
	fmt.Println(myString())
	fmt.Println(myString())
	fmt.Println(myString())
	fmt.Println(myString())
	fmt.Println(myString())
}
