package main

import (
	"fmt"
	"sync"
	"time"
)

func PrintFrom1To5(number string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i < 6; i++ {
		time.Sleep(200 * time.Millisecond)
		fmt.Println(number, i)
	}
}
func main() {
	fmt.Println("START")

	wg := sync.WaitGroup{}
	wg.Add(3)

	go PrintFrom1To5("first", &wg)
	go PrintFrom1To5("second", &wg)
	go PrintFrom1To5("Three", &wg)

	wg.Wait()

	fmt.Println("DONE")
}
