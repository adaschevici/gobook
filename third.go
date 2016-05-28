package main

import (
	"os"
)

func main() {
	file, _ := os.Create("defer.txt")
	defer file.Close()

	//	for {
	//		break
	//	}
}
