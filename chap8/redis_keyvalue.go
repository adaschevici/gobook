package main

import (
	"fmt"
	"log"

	"github.com/garyburd/redigo/redis"
)

func main() {

	connection, _ := redis.Dial("tcp", ":6379")
	defer connection.Close()

	// data, err := redis.Values(connection.Do("SORT", "Users", "BY", "User:*->name",
	// 	"GET", "User:*->name"))

	// if err != nil {
	// 	fmt.Println("Error getting values", err)
	// }

	// spew.Dump(data)
	// for _ = range data {
	// 	var Uname string
	// 	data, err = redis.Scan(data, &Uname)
	// 	if err != nil {
	// 		fmt.Println("Error getting value", err)
	// 	} else {
	// 		fmt.Println("Name Uname")
	// 	}
	// }
	data, err := connection.Do("INFO")

	if err != nil {
		fmt.Println("Error connecting to redis.")
	} else {
		log.Printf("info=%s", data)
	}

	inserted, _ := connection.Do("SET", "awesome", "I inserted some crap")
	log.Printf("Inserted=%s", inserted)

	getme, _ := connection.Do("GET", "awesome")
	log.Printf("Got: %s", getme)

}
