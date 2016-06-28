package main

import (
	"errors"
	"log"
	"reflect"
)

// func Add(x int, y int) int {
// 	return x + y
// }

func Add(x int, y int) (int, error) {
	var err error

	xType := reflect.TypeOf(x).Kind()
	yType := reflect.TypeOf(y).Kind()
	if xType != reflect.Int || yType != reflect.Int {
		log.Printf("Type of x is %v", xType)
		err = errors.New("Incorrect type for integer a or b!")
	}
	return x + y, err
}

func main() {
	e, v := Add(3, 1)
	log.Printf("%v %v", e, v)
	log.Printf("Try me now will ya?")
}
