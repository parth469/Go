package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Worker(readerStream <-chan string, done <-chan interface{}) {
	for {
		select {
		case v := <-readerStream:
			time.Sleep(499 * time.Millisecond)
			fmt.Println("value", v)
		case <-done:
			fmt.Println("Parant close this channal")
			return
		}
	}
}

func Task(done chan<- interface{}) <-chan string {
	task := make(chan string)

	go func() {
		defer close(task)
		defer close(done)

		for i := range 10 {
			if i == 5 {
				return
			}
			task <- fmt.Sprintln("work on", i)
		}
	}()

	return task
}
func main() {
	// doWork := func(
	// 	done <-chan interface{},
	// 	strings <-chan string,
	// ) <-chan interface{} {
	// 	terminated := make(chan interface{})
	// 	go func() {
	// 		defer fmt.Println("doWork exited.")
	// 		defer close(terminated)
	// 		for {
	// 			select {
	// 			case s := <-strings:
	// 				fmt.Println(s)
	// 			case <-done:
	// 				return
	// 			}
	// 		}
	// 	}()
	// 	return terminated
	// }
	// done := make(chan interface{})
	// terminated := doWork(done, nil)
	// go func() {
	// 	// Cancel the operation after 1 second.
	// 	time.Sleep(1 * time.Second)
	// 	fmt.Println("Canceling doWork goroutine...")
	// 	close(done)
	// }()
	// <-terminated
	// fmt.Println("Done.")

	newRandStream := func(done <-chan interface{}) <-chan int {
		randStream := make(chan int)
		go func() {
			defer fmt.Println("newRandStream closure exited.")
			defer close(randStream)
			for {
				select {
				case <-done:
					return
				case randStream <- rand.Int():
				}
			}
		}()
		return randStream
	}

	done := make(chan interface{})

	randStream := newRandStream(done)
	fmt.Println("3 random ints:")

	for i := 1; i <= 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}

	close(done)
	// Simulate ongoing work
	time.Sleep(1 * time.Second)
}
