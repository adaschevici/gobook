package main

import (
	"log"
	"os"
	"strconv"
	"sync"
)

const totalGoroutines = 5

type Worker struct {
	wLog *log.Logger
	Name string
}

func main() {
	// done := make(chan bool)
	var wg sync.WaitGroup

	wg.Add(totalGoroutines)
	for i := 0; i < totalGoroutines; i++ {

		myWorker := Worker{}
		myWorker.Name = "Goroutine " + strconv.FormatInt(int64(i), 10) + ""
		myWorker.wLog = log.New(os.Stdout, myWorker.Name, 1)
		go func(w *Worker) {
			defer wg.Done()
			w.wLog.Print("Hmm")
			// done <- true
		}(&myWorker)
	}
	log.Println("...")
	wg.Wait()
	// <-done
}
