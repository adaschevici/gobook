package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
)

func main() {
	someUser, _ := user.Current()
	currentUser := someUser.Username
	var listenFolder = fmt.Sprintf("/Users/%s/couchbase_volume/", currentUser)
	logFile, _ := os.OpenFile(fmt.Sprintf("%stest.log", listenFolder), os.O_RDWR, 0755)

	log.SetOutput(logFile)
	log.Println("Sending an entry to log!")

	logFile.Close()
}
