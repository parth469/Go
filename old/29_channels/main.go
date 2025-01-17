package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(3)
	go func() {
		defer wg.Done()
		fmt.Println("Hello Welcome")
		ch <- 5
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Hello Welcome")
		fmt.Println(<-ch) // Receive the first value from the channel
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Hello Welcome")
		ch <- 15 // Send the second value to the channel
	}()
	wg.Wait()
	close(ch)             // Close the channel to avoid deadlock
	for val := range ch { // Safely receive any remaining values
		fmt.Println(val, "Done")
	}
}
