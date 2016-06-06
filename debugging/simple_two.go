package main

import "fmt"

type pair struct {
	x int
	y int
}

func handleNumber(i int) *pair {
	val := i
	if i%2 == 0 {
		val = f(i)
	}
	return &pair{
		x: i,
		y: val,
	}
}

func f(x int) int {
	return x*x + x
}

func main() {
	for i := 0; i < 5; i++ {
		p := handleNumber(i)
		fmt.Printf("%v\n", p)
		fmt.Println("looping")
	}

}
