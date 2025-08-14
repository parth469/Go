package main

import (
	"fmt"
	"time"
)

func main() {
	// start := time.Now()

	// ch := make(chan int)
	// ch1 := make(chan int)

	// go func() {
	// 	defer close(ch)
	// 	time.Sleep(5 * time.Second)
	// 	ch <- 10
	// }()

	// go func() {
	// 	defer close(ch1)
	// 	time.Sleep(5 * time.Second)
	// 	ch1 <- 9
	// }()

	// select {
	// case v := <-ch:
	// 	fmt.Println("value is ", v, time.Since(start))
	// case v := <-ch1:
	// 	fmt.Println("value is ", v, time.Since(start))
	// case <-time.After(4 * time.Second):
	// 	fmt.Println("time out")
	// }

	done := make(chan interface{})
	go func() {
		time.Sleep(5 * time.Second)
		close(done)
	}()
	workCounter := 0
loop:
	for {
		select {
		case <-done:
			break loop
		default:
		}
		// Simulate work
		workCounter++
		time.Sleep(1 * time.Second)
	}
	fmt.Printf("Achieved %v cycles of work before signalled to stop.\n", workCounter)
}
