package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func SayHi() {
	fmt.Println("Hello")
}

func BenchmarkContextSwitch(b *testing.B) {
	var wg sync.WaitGroup
	begin := make(chan struct{})
	c := make(chan struct{})
	var token struct{}
	sender := func() {
		defer wg.Done()
		<-begin
		for i := 0; i < b.N; i++ {
			c <- token
		}
	}
	receiver := func() {
		defer wg.Done()
		<-begin
		for i := 0; i < b.N; i++ {
			<-c
		}
	}
	wg.Add(2)
	go sender()
	go receiver()
	b.StartTimer()
	close(begin)
	wg.Wait()
}

func main() {
	var wg sync.WaitGroup
	say := "hello"
	wg.Add(1)
	// first-class functions ( can able to assing function to var)
	Anonymouse := func() {
		say = "DONE"
		fmt.Println("Anonymouse Assign")
		wg.Done()

	}

	go SayHi()

	func() {
		fmt.Println("Anonymouse Call")
	}()

	Anonymouse()

	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(salutation)
		}()
	}

	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func(salutat string) {
			defer wg.Done()
			fmt.Println(salutat)
		}(salutation)
	}
	wg.Wait()

	fmt.Println(say)

	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	orderReady := false
	customerNumber := 0

	// Customers
	for i := 1; i <= 3; i++ {
		go func(id int) {
			mu.Lock()
			for !orderReady || customerNumber != id {
				fmt.Printf("Customer %d is waiting for their turn...\n", id)
				cond.Wait()
			}
			fmt.Printf("Customer %d got their food!\n", id)
			orderReady = false // Reset so next plate can be prepared
			mu.Unlock()
		}(i)
	}

	// Chef prepares one plate at a time
	for i := 1; i <= 3; i++ {
		time.Sleep(2 * time.Second) // Cooking time
		mu.Lock()
		customerNumber = i
		orderReady = true
		fmt.Printf("Chef: Plate for Customer %d is ready!\n", i)
		cond.Signal() // Wake exactly one waiting customer
		mu.Unlock()
	}

	time.Sleep(2 * time.Second) // Wait for all to finish

}
