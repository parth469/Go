package main

import (
	"fmt"
	"sync"
)

func increaseByone(counter *int, lock *sync.Mutex, wg *sync.WaitGroup) {
	lock.Lock()
	*counter = *counter + 1
	lock.Unlock()
	wg.Done()

}
func main() {
	lock := sync.Mutex{}
	wg := sync.WaitGroup{}
	counter := 0
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go increaseByone(&counter, &lock, &wg)
	}
	wg.Wait()

	fmt.Println("total Count", counter)

}
