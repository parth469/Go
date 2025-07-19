package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)

	wg := sync.WaitGroup{}

	wg.Add(2)
	go func() {
		defer close(ch)
		defer wg.Done()

		for i := 1; i < 11; i++ {
			ch <- i
		}
	}()

	go func() {
		for value := range ch {
			fmt.Println(value)
		}
		wg.Done()
	}()
	wg.Wait()

}
