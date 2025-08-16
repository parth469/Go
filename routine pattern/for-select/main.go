package main

import (
	"fmt"
	"time"
)

type Item struct {
	Index int
	Value rune
}

func main() {
	ch := make(chan Item)

	go func() {
		message := []rune{'a', 'b', 'c', 'd'}
		for i, v := range message {
			ch <- Item{i, v}
		}
		close(ch)
	}()

	for item := range ch {
		fmt.Printf("Index: %d, Value: %c\n", item.Index, item.Value)
	}
	loopOver()
}

func loopOver() {
	done := make(chan struct{}, 1)

	for {
		select {
		case <-done:
			fmt.Println("Done")
			return // exit the loop safely
		default:
			fmt.Println("start Process")
			time.Sleep(2 * time.Second)
			fmt.Println("Task done")
			done <- struct{}{} // safe because channel isn't closed yet
		}
	}
}