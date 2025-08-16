package main

import (
	"fmt"
	"time"
)

func Start(records []int) <-chan int {
	processChan := make(chan int)

	go func() {
		defer close(processChan)
		for _, v := range records {
			data := v
			fmt.Println("START ", data)
			time.Sleep(100 * time.Millisecond)
			processChan <- data
		}
	}()

	return processChan
}
func Add(ch <-chan int, addValue int) <-chan int {
	processChan := make(chan int)

	go func() {
		defer close(processChan)
		for v := range ch {
			data := v + addValue
			fmt.Println("Process", data)
			time.Sleep(100 * time.Millisecond)
			processChan <- data
		}
	}()
	return processChan
}

func Done(ch <-chan int) <-chan interface{} {
	doneCh := make(chan interface{})
	go func() {
		defer close(doneCh)
		for v := range ch {
			fmt.Println("Done", v)
		}
	}()

	return doneCh
}

func main() {
	record := []int{1, 2, 3, 4, 65, 6}
	fmt.Println(record)
	procesStream := Start(record)
	doneStream := Add(procesStream, 5)
	<-Done(doneStream)

	fmt.Println("Done All procress")
}
