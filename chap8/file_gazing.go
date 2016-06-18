package main

import (
	"fmt"
	"log"
	"os/user"

	"github.com/howeyc/fsnotify"
)

func main() {
	scriptDone := make(chan bool)
	dirSpy, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	someUser, _ := user.Current()
	currentUser := someUser.Username

	go func() {
		for {
			select {
			case fileChange := <-dirSpy.Event:
				log.Println("Something happened to a file", fileChange)
			case err := <-dirSpy.Error:
				log.Println("Error with fsnotify:", err)
			}
		}
	}()

	err = dirSpy.Watch(fmt.Sprintf("/Users/%s/go_play/demo/chap8/", currentUser))
	if err != nil {
		fmt.Println(err)
	}

	<-scriptDone

	dirSpy.Close()
}
