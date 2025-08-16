package main

import "fmt"

func Producer(sender chan<- int) {
	defer close(sender)

	for i := range 10 {
		sender <- i + 1
	}
}
func main() {
	senderChan := make(chan int)

	go Producer(senderChan)

	for v := range senderChan {
		fmt.Println(v)
	}
}
