package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	REQUEST_LIMIT = 2
	API_CALL      = 10
	REQUEST_RATE  = 200 * time.Millisecond
	PROCESS_TIME  = 600 * time.Millisecond
)

func main() {
	req := make(chan int, REQUEST_LIMIT)
	wg := sync.WaitGroup{}

	// Producer
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= API_CALL; i++ {
			select {
			case req <- i:
				fmt.Printf("%d\n", i)
			default:
				fmt.Println("Request limit reached, Drop")
				time.Sleep(REQUEST_RATE)
			}
			time.Sleep(REQUEST_RATE)
		}
		close(req)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for id := range req {
			fmt.Printf("Processing request %d\n", id)
			time.Sleep(PROCESS_TIME)
			fmt.Printf("Completed request %d\n", id)
		}
	}()

	wg.Wait()
	fmt.Println("All requests processed")
}
