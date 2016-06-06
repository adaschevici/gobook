package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

// TODO: this is still racy according to the golang race detector
var balance int
var transactionNo int
var mutex sync.Mutex

func main() {
	rand.Seed(time.Now().Unix())
	runtime.GOMAXPROCS(2)
	var wg sync.WaitGroup
	balanceChan := make(chan int)
	tranChan := make(chan bool)

	balance = 1000
	transactionNo = 0
	fmt.Println("Starting balance: $", balance)

	wg.Add(1)
	for i := 0; i < 100; i++ {
		go func(ii int) {
			transactionAmount := rand.Intn(25)
			balanceChan <- transactionAmount

			if ii == 99 {
				fmt.Println("We should be quitting now")
				tranChan <- true
				close(balanceChan)
				wg.Done()
			}
		}(i)
	}

	// TODO: this is where the race condition still happens
	go transaction(0)
	breakpoint := false
	for {
		if breakpoint == true {
			break
		}
		select {
		case amt := <-balanceChan:
			fmt.Println("Transaction for $", amt)
			cond := balance - amt
			if (cond) < 0 {
				fmt.Println("Transaction failed")
			} else {
				balance = cond
				fmt.Println("Transaction succeeded")
			}
			fmt.Println("Balance now $", balance)
		case status := <-tranChan:
			if status == true {
				fmt.Println("Done")
				breakpoint = true
				close(tranChan)
			}
		}
	}
	wg.Wait()

	fmt.Println("Final balance: $", balance)
}

func transaction(amt int) bool {
	mutex.Lock()
	defer mutex.Unlock()
	approved := false
	if (balance - amt) < 0 {
		approved = false
	} else {
		approved = true
		balance = balance - amt
	}

	approvedText := "declined"
	if approved == true {
		approvedText = "approved"
	}
	transactionNo = transactionNo + 1
	fmt.Println(transactionNo, "Transaction for $", amt, approvedText)
	fmt.Println("\tRemaining balance $", balance)
	return approved
}
