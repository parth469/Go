package main

import (
	"fmt"
	"sync"
)

var cond = sync.NewCond(&sync.Mutex{})
var queue []int

const maxQueueSize = 2

func addToQueue(task int) {
	cond.L.Lock()
	defer cond.L.Unlock()

	// Wait until there's space in the queue
	for len(queue) >= maxQueueSize {
		fmt.Println("Queue full, waiting to add:", task)
		cond.Wait()
	}

	// Add the task
	queue = append(queue, task)
	fmt.Println("Added:", task, "Queue:", queue)

	// Wake up one waiting remover
	cond.Signal()
}

func removeFromQueue() {
	cond.L.Lock()
	defer cond.L.Unlock()

	// Wait until queue is not empty
	for len(queue) == 0 {
		fmt.Println("Queue empty, waiting to remove")
		cond.Wait()
	}

	// Remove the first item
	removed := queue[0]
	queue = queue[1:]
	fmt.Println("Removed:", removed, "Queue:", queue)

	// Wake up one waiting adder
	cond.Signal()
}

func main() {
	type Button struct {
		Clicked *sync.Cond
	}
	button := Button{Clicked: sync.NewCond(&sync.Mutex{})}
	subscribe := func(c *sync.Cond, fn func()) {
		var goroutineRunning sync.WaitGroup
		goroutineRunning.Add(1)
		go func() {
			goroutineRunning.Done()
			c.L.Lock()
			defer c.L.Unlock()
			c.Wait()
			fn()
		}()
		goroutineRunning.Wait()
	}
	var clickRegistered sync.WaitGroup
	clickRegistered.Add(1)
	subscribe(button.Clicked, func() {
		fmt.Println("Maximizing window.")
		clickRegistered.Done()
	})
	subscribe(button.Clicked, func() {
		fmt.Println("Displaying annoying dialog box!")
		clickRegistered.Done()
	})
	subscribe(button.Clicked, func() {
		fmt.Println("Mouse clicked.")
		clickRegistered.Done()
	})
	button.Clicked.Signal()
	clickRegistered.Wait()

	var coutner int
	wg := sync.WaitGroup{}
	ones := sync.Once{}
	increment := func() {
		coutner++
	}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			ones.Do(increment)
		}()
	}
	wg.Wait()

	fmt.Println(coutner)

	var a, b sync.Once

	var initB func()

	intiA := func() {
		b.Do(initB)
	}
	initB = func() { a.Do(intiA) }
	a.Do(intiA)
}
