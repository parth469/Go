package main

import (
	"fmt"
	"time"
)

func sayHelo() {
	fmt.Println("Hello")
}

func sayHi() {
	fmt.Println("sayHi")
}
func main() {
	fmt.Println("Start")
	go sayHelo()
	go sayHi()

	time.Sleep(time.Millisecond)
	fmt.Println("Done")
}
