package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"sync"
	"time"
)

var applicationStatus bool
var urls []string
var fullText string
var urlsProcessed int
var foundUrls []string
var pageContent string
var totalURLCount int
var wg sync.WaitGroup

func readURLs(statusChannel chan int, textChannel chan string) {

	time.Sleep(time.Millisecond * 1)
	fmt.Println("Grabbing", len(urls), "urls")
	for i := 0; i < totalURLCount; i++ {

		fmt.Println("Url", i, urls[i])
		fmt.Println("getting url", urls[i])
		resp, _ := http.Get(urls[i])
		text, err := ioutil.ReadAll(resp.Body)

		fmt.Println(string(text))
		textChannel <- string(text)

		if err != nil {
			fmt.Println("HTML body is not retreived")
		}

		statusChannel <- 0
	}
}

func addToScrapedText(textChannel chan string, processChannel chan bool) {
	for {
		select {
		case pC := <-processChannel:
			if pC == true {
				// hang on
			}
			if pC == false {
				close(textChannel)
				close(processChannel)
			}
		case tC := <-textChannel:
			fmt.Println("Adding to full text")
			fullText += tC
			fmt.Println(fullText)
		}
	}
}

func evaluateStatus(statusChannel chan int, textChannel chan string, processChannel chan bool) {
	for {
		select {
		case status := <-statusChannel:
			fmt.Println("URLS Processed:", urlsProcessed, "VS total urls count:", totalURLCount)
			urlsProcessed++
			if status == 0 {
				fmt.Println("Urls processed: ", urlsProcessed)
			}
			if status == 1 {
				fmt.Println("Closing status channel")
				close(statusChannel)
			}
			if urlsProcessed == totalURLCount {
				fmt.Println("Read all top level URLs")
				processChannel <- false
				applicationStatus = false
			}
		}
	}
}

func main() {
	runtime.GOMAXPROCS(2)
	applicationStatus = true
	statusChannel := make(chan int)
	textChannel := make(chan string)
	processChannel := make(chan bool)
	totalURLCount = 0

	urls = append(urls, "http://localhost:8080/")
	urls = append(urls, "http://localhost:8080/")
	urls = append(urls, "http://localhost:8080/")
	urls = append(urls, "http://localhost:8080/")
	urls = append(urls, "http://localhost:8080/")

	fmt.Println("Starting spider")

	urlsProcessed = 0
	totalURLCount = len(urls)

	go evaluateStatus(statusChannel, textChannel, processChannel)

	go readURLs(statusChannel, textChannel)

	go addToScrapedText(textChannel, processChannel)

	for {
		if applicationStatus == false {
			fmt.Println(fullText)
			fmt.Println("Done!")
			break
		}
		select {
		case sC := <-statusChannel:
			fmt.Println("Message on StatusChannel", sC)
		}
	}
}
