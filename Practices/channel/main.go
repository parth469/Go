package main

import (
	"fmt"
	"time"
)

// func main() {
// 	intStream := make(chan int, 4)
// 	go func() {
// 		defer close(intStream)
// 		fmt.Println("Producer starting...")
// 		for i := 0; i < 5; i++ {
// 			fmt.Printf("Sending: %d\n", i)
// 			intStream <- i
// 		}
// 		fmt.Println("Producer Done.")
// 	}()
// 	for integer := range intStream {
// 		fmt.Printf("Received %v.\n", integer)
// 	}
// }

func Producer(writerChan *chan<- int, dataStream *chan int) {
	defer close(*dataStream)
	defer fmt.Println("Proucer Done")

	for i := 0; i < 10; i++ {
		fmt.Println("New Request ", i)
		*writerChan <- i
		time.Sleep(time.Duration(100*i) * time.Millisecond)
	}
}

func Consumer(readerChan *<-chan int) {
	defer fmt.Println("All Message consume Done")

}

func main() {

	dataStream := make(chan int)
	var readerChan <-chan int
	var writerChan chan<- int

	readerChan = dataStream
	writerChan = dataStream

	go Producer(&writerChan, &dataStream)

	for value := range readerChan {
		fmt.Println("Process Done for", value)
	}
	// Create an unbuffered channel of type int
	// ch := make(chan int)
	// Create a buffered channel of type int with capacity 1
	// chBuff := make(chan int, 1)
	// closeCh := make(chan int)
	// var closeChNil chan int

	// if closeCh == nil {
	// 	fmt.Println("It nill chane")
	// }

	// if closeChNil == nil {
	// 	fmt.Println("It nill chane")
	// }

	// close(closeCh)

	// // Attempt to receive from the buffered channel (will block if empty)
	// fmt.Println(<-chBuff, "BUFF channel")
	// // Attempt to receive from the unbuffered channel (will block forever)
	// fmt.Println(<-ch, "panic")
}
