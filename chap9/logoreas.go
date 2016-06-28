package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"strconv"
)

var totalGoroutines = 5

type Worker struct {
	FileName string
	File     *os.File
	wLog     *log.Logger
	Name     string
}

func main() {
	done := make(chan bool)
	someUser, _ := user.Current()
	currentUser := someUser.Username
	var listenFolder = fmt.Sprintf("/Users/%s/couchbase_volume/", currentUser)
	logFile, _ := os.OpenFile(fmt.Sprintf("%stest.log", listenFolder), os.O_RDWR, 0755)
	for i := 0; i < totalGoroutines; i++ {

		myWorker := Worker{}
		myWorker.Name = "Goroutine " + strconv.FormatInt(int64(i), 10) + ""
		myWorker.FileName = listenFolder + strconv.FormatInt(int64(i), 10) + ".log"
		tmpFile, _ := os.OpenFile(myWorker.FileName, os.O_CREATE, 0755)
		myWorker.File = tmpFile
		myWorker.wLog = log.New(myWorker.File, myWorker.Name, 1)
		go func(w *Worker) {

			w.wLog.Print("Hmm")

			done <- true
		}(&myWorker)
	}

	log.SetOutput(logFile)
	log.Println("Sending an entry to log!")

	logFile.Close()
}
