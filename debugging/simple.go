package main

// This is for checking out debugging using gdb
import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("looping")
	}
	fmt.Println("Done")
}
