package main

import (
	"fmt"
	"runtime"
)

func showNumber(num int) {
	fmt.Println(num)
}

func main() {
	iterations := 10
	runtime.GOMAXPROCS(2)
	// This yields control from goroutines to finish
	// opposed to syn waitgroup where you have to register
	// a number of goroutines.
	for i := 0; i <= iterations; i++ {

		go showNumber(i)

	}
	fmt.Println("Goodbye!")
	runtime.Gosched()

}
