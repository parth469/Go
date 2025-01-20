package main

import "fmt"

func main() {
	ch := make(chan int, 3)
	ch <- 5
	fmt.Printf("Len: %d\n", len(ch))

	ch <- 6
	fmt.Printf("Len: %d\n", len(ch))
	ch <- 7
	fmt.Printf("Len: %d\n", len(ch))
	close(ch)
	for v := range ch {
		fmt.Println(v)
	}
}
