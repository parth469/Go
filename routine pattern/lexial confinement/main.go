package main

import (
	"bytes"
	"fmt"
	"sync"
)

func Producer() <-chan int {
	sender := make(chan int)
	go func() {
		defer close(sender)
		for i := 0; i < 10; i++ {
			sender <- i + 1
		}
	}()

	return sender
}
func Consumer(reciver <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range reciver {
		fmt.Println(v)
	}
}
func main() {

	wg := sync.WaitGroup{}
	wg.Add(1)

	go Consumer(Producer(), &wg)

	wg.Wait()

	var wg1 sync.WaitGroup
	wg1.Add(2)
	data := []byte("golang")
	go PrintData(&wg1, data[:3])
	go PrintData(&wg1, data[3:])

	wg1.Wait()
}

func PrintData(wg *sync.WaitGroup, data []byte) {
	defer wg.Done()
	var buff bytes.Buffer
	for _, b := range data {
		fmt.Fprintf(&buff, "%c", b)
	}
	fmt.Println(buff.String())
}
