package main

import "fmt"

func main() {

	aValue := new(int)

	defer fmt.Println(*aValue)

	for i := 0; i < 100; i++ {
		*aValue++
	}
}
