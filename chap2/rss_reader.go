package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"time"

	"github.com/ajstarks/svgo"
)

var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func Init(
	traceHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {

	Trace = log.New(traceHandle,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

type FeedItem struct {
	feedIndex int
	complete  bool
	url       string
}

type Feed struct {
	url           string
	status        int
	itemCount     int
	complete      bool
	itemsComplete bool
	index         int
}

var feeds []Feed
var height int
var width int
var colors string
var startTime int64
var timeout int
var feedSpace int
var wg sync.WaitGroup

func grabFeed(feed *Feed, feedChan chan bool, osvg *svg.SVG) {
	startGrab := time.Now().Unix()
	startGrabSeconds := startGrab - startTime

	fmt.Println("Grabbing feed", feed.url, "at", startGrabSeconds, "second mark")

	if feed.status == 0 {
		fmt.Println("Feed not read yet")
		feed.status = 1
	}
}

func main() {
	Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
}
