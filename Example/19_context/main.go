package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	done := make(chan string, 2)

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Done First")
		done <- "first"
	}()

	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("Done Second")
		done <- "second"
	}()

	select {
	case <-ctx.Done():
		fmt.Println("Complete")
	case <-done:
		fmt.Println("Goroutine finished before timeout")
	}
}
