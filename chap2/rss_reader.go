package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"

	"github.com/ajstarks/svgo"
	rss "github.com/jteeuwen/go-pkg-rss"
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
var colors []string
var startTime int64
var timeout int
var feedSpace int
var wg sync.WaitGroup

func grabFeed(feed *Feed, feedChan chan bool, osvg *svg.SVG) {
	startGrab := time.Now().Unix()
	startGrabSeconds := startGrab - startTime

	Info.Println("Grabbing feed", feed.url, "at", startGrabSeconds, "second mark")

	if feed.status == 0 {
		Info.Println("Feed not read yet")
		feed.status = 1

		startX := int(startGrabSeconds * 33)
		startY := feedSpace * (feed.index)

		Info.Println(startY)
		wg.Add(1)

		rssFeed := rss.New(timeout, true, channelHandler, itemsHandler)

		if err := rssFeed.Fetch(feed.url, nil); err != nil {
			Info.Println(os.Stderr, "[e] %s: %s", feed.url, err)
			return
		} else {
			endSec := time.Now().Unix()
			endX := int((endSec - startGrab))
			if endX == 0 {
				endX = 1
			}
			Info.Println("Read feed in", endX, "seconds")
			osvg.Rect(startX, startY, endX, feedSpace, "fill:#000000;opacity:.4")
			wg.Wait()
			endGrab := time.Now().Unix()
			endGrabSeconds := endGrab - startTime
			feedEndX := int(endGrabSeconds * 33)

			osvg.Rect(feedEndX, startY, 1, feedSpace, "fill:#ff0000;opacity:.9")

			feedChan <- true
		}
	} else if feed.status == 1 {
		Info.Println("Feed already in progress")
	}
}

func channelHandler(feed *rss.Feed, newchannels []*rss.Channel) {

}

func itemsHandler(feed *rss.Feed, ch *rss.Channel, newitems []*rss.Item) {

	Info.Println("Found", len(newitems), "items in", feed.Url)

	for i := range newitems {
		url := *newitems[i].Guid
		Info.Println(url)

	}

	wg.Done()
}

func getRSS(rw http.ResponseWriter, req *http.Request) {
	startTime = time.Now().Unix()
	rw.Header().Set("Content-Type", "image/svg+xml")
	outputSVG := svg.New(rw)
	outputSVG.Start(width, height)

	feedSpace = (height - 20) / len(feeds)

	for i := 0; i < 30000; i++ {
		timeText := strconv.FormatInt(int64(i/10), 10)
		if i%1000 == 0 {
			outputSVG.Text(i/30, 390, timeText, "text-anchor:middle;font-size:10px;fill:#000000")
		} else if i%4 == 0 {
			outputSVG.Circle(i, 377, 1, "fill:#cccccc;stroke:none")
		}

		if i%10 == 0 {
			outputSVG.Rect(i, 0, 1, 400, "fill:#dddddd")
		}
		if i%50 == 0 {
			outputSVG.Rect(i, 0, 1, 400, "fill:#cccccc")
		}

	}

	feedChan := make(chan bool, 2)

	for i := range feeds {

		outputSVG.Rect(0, (i * feedSpace), width, feedSpace, "fill:"+colors[i]+";stroke:none;")
		feeds[i].status = 0
		go grabFeed(&feeds[i], feedChan, outputSVG)
		result := <-feedChan
		Info.Println(result)

	}

	outputSVG.End()
}

func main() {
	Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	runtime.GOMAXPROCS(2)

	timeout = 1000

	width = 1000
	height = 400

	feeds = append(feeds, Feed{index: 0, url: "https://groups.google.com/forum/feed/golang-nuts/msgs/rss_v2_0.xml?num=50", status: 0, itemCount: 0, complete: false, itemsComplete: false})
	feeds = append(feeds, Feed{index: 1, url: "https://groups.google.com/forum/feed/golang-nuts/msgs/rss_v2_0.xml?num=50", status: 0, itemCount: 0, complete: false, itemsComplete: false})
	// feeds = append(feeds, Feed{index: 1, url: "http://www.reddit.com/r/golang/.rss", status: 0, itemCount: 0, complete: false, itemsComplete: false})
	feeds = append(feeds, Feed{index: 2, url: "https://groups.google.com/forum/feed/golang-dev/msgs/rss_v2_0.xml?num=50", status: 0, itemCount: 0, complete: false, itemsComplete: false})

	colors = append(colors, "#ff9999")
	colors = append(colors, "#99ff99")
	colors = append(colors, "#9999ff")

	http.Handle("/getrss", http.HandlerFunc(getRSS))
	err := http.ListenAndServe(":1900", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}
