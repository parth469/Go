package main

import (
	"fmt"
	"sync"
	"time"
)

var counter int = 9
var lock sync.RWMutex

func main() {
	wg := sync.WaitGroup{}
	hello := func(wg *sync.WaitGroup, i int) {
		defer wg.Done()
		fmt.Println("hello for", i)
		lock.Lock()
		if i%2 == 0 {
			counter += 1
		}
		time.Sleep(2 * time.Millisecond)
		fmt.Println(counter, "i", i)
		lock.Unlock()
	}

	numberOFloop := 5

	wg.Add(numberOFloop)
	i := 0
	for i < numberOFloop {
		i++
		go hello(&wg, i)
	}

	wg.Wait()
	fmt.Println(counter)
}
