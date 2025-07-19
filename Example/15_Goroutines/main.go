package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	fmt.Println("Go Concurrency Patterns Deep Dive")

	// Basic patterns
	demoBasicGoroutine()
	demoConcurrentCounter()
	demoChannelBasics()

	// Intermediate patterns
	demoWorkerPool()
	demoSelectStatement()
	demoFanOutFanIn()

	// Advanced patterns
	demoContextCancellation()
	demoErrorHandling()
	demoRateLimiting()
	demoAtomicOperations()
	demoOnceUsage()
	demoPipelinePattern()
	demoPubSubPattern()
}

// 1. Basic Goroutine with WaitGroup
func demoBasicGoroutine() {
	fmt.Println("\n--- 1. Basic Goroutine Example ---")

	var wg sync.WaitGroup

	// Launch 3 goroutines
	wg.Add(3)

	for i := 0; i < 3; i++ {
		go func(id int) {
			defer wg.Done()
			time.Sleep(time.Duration(id) * 100 * time.Millisecond)
			fmt.Printf("Goroutine %d completed\n", id)
		}(i)
	}

	wg.Wait()
	fmt.Println("All basic goroutines completed")
}

// 2. Concurrent Counter with Mutex
func demoConcurrentCounter() {
	fmt.Println("\n--- 2. Concurrent Counter Example ---")

	var (
		counter int
		lock    sync.Mutex
		wg      sync.WaitGroup
	)

	const numGoroutines = 1000
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			lock.Lock()
			defer lock.Unlock()
			defer wg.Done()
			counter++
		}()
	}

	wg.Wait()
	fmt.Println("Final counter value:", counter)
}

// 3. Channel Basics
func demoChannelBasics() {
	fmt.Println("\n--- 3. Channel Basics ---")

	// Unbuffered channel (synchronous)
	messageChan := make(chan string)

	go func() {
		time.Sleep(500 * time.Millisecond)
		messageChan <- "ello from goroutine!"
	}()

	// Main goroutine blocks until message is received
	msg := <-messageChan
	fmt.Println("Received:", msg)

	// Buffered channel (asynchronous)
	bufferedChan := make(chan int, 3)

	// Can send multiple values without blocking
	bufferedChan <- 1
	bufferedChan <- 2
	bufferedChan <- 3

	// Close channel when no more values will be sent
	close(bufferedChan)

	// Range over channel to receive values
	fmt.Println("Buffered channel values:")
	for val := range bufferedChan {
		fmt.Println(val)
	}
}

// 4. Worker Pool Pattern
func demoWorkerPool() {
	fmt.Println("\n--- 4. Worker Pool Pattern ---")

	const (
		numJobs    = 10
		numWorkers = 3
	)

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	var wg sync.WaitGroup

	// Start workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// Send jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Wait for all workers to complete
	wg.Wait()
	close(results)

	// Collect results
	fmt.Println("Job results:")
	for result := range results {
		fmt.Println(result)
	}
}

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(500 * time.Millisecond) // Simulate work
		results <- job * 2                 // Return processed result
	}
}

// 5. Select Statement with Channels
func demoSelectStatement() {
	fmt.Println("\n--- 5. Select Statement ---")

	chan1 := make(chan string)
	chan2 := make(chan string)
	quit := make(chan bool)

	// Start goroutines that will send to channels
	go func() {
		time.Sleep(800 * time.Millisecond)
		chan1 <- "from channel 1"
	}()

	go func() {
		time.Sleep(300 * time.Millisecond)
		chan2 <- "from channel 2"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		quit <- true
	}()

	// Select waits on multiple channel operations
	for {
		select {
		case msg1 := <-chan1:
			fmt.Println("Received:", msg1)
		case msg2 := <-chan2:
			fmt.Println("Received:", msg2)
		case <-quit:
			fmt.Println("Quitting")
			return
		case <-time.After(500 * time.Millisecond):
			fmt.Println("Timeout waiting for channels")
		}
	}
}

// 6. Fan-out/Fan-in Pattern
func demoFanOutFanIn() {
	fmt.Println("\n--- 6. Fan-out/Fan-in Pattern ---")

	// Generate numbers
	numbers := make(chan int)
	go func() {
		for i := 1; i <= 10; i++ {
			numbers <- i
		}
		close(numbers)
	}()

	// Fan-out: Distribute work across multiple goroutines
	const numWorkers = 3
	results := make(chan int)
	var wg sync.WaitGroup

	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go func(workerID int) {
			defer wg.Done()
			for n := range numbers {
				// Simulate work
				time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
				fmt.Printf("Worker %d processed %d\n", workerID, n)
				results <- n * n
			}
		}(i)
	}

	// Fan-in: Collect results in a separate goroutine
	go func() {
		wg.Wait()
		close(results)
	}()

	// Process results
	sum := 0
	for res := range results {
		sum += res
	}
	fmt.Printf("Sum of squares: %d\n", sum)
}

// 7. Context Cancellation
func demoContextCancellation() {
	fmt.Println("\n--- 7. Context Cancellation ---")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Cancelling operation...")
		cancel()
	}()

	select {
	case <-time.After(2 * time.Second):
		fmt.Println("Operation completed successfully")
	case <-ctx.Done():
		fmt.Println("Operation cancelled:", ctx.Err())
	}
}

// 8. Error Handling in Goroutines
func demoErrorHandling() {
	fmt.Println("\n--- 8. Error Handling ---")

	errChan := make(chan error)
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		time.Sleep(500 * time.Millisecond)
		errChan <- nil // No error
	}()

	go func() {
		defer wg.Done()
		time.Sleep(300 * time.Millisecond)
		errChan <- errors.New("something went wrong")
	}()

	go func() {
		wg.Wait()
		close(errChan)
	}()

	for err := range errChan {
		if err != nil {
			fmt.Println("Error received:", err)
		} else {
			fmt.Println("Task completed successfully")
		}
	}
}

// 9. Rate Limiting
func demoRateLimiting() {
	fmt.Println("\n--- 9. Rate Limiting ---")

	requests := make(chan int, 10)
	for i := 1; i <= 10; i++ {
		requests <- i
	}
	close(requests)

	// Limiter channel will receive a value every 200ms
	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		<-limiter // Block until we can proceed
		fmt.Println("Processing request", req, time.Now())
	}
}

// 10. Atomic Operations
func demoAtomicOperations() {
	fmt.Println("\n--- 10. Atomic Operations ---")

	var counter int64
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1)
		}()
	}

	wg.Wait()
	fmt.Println("Atomic counter value:", counter)
}

// 11. sync.Once Usage
func demoOnceUsage() {
	fmt.Println("\n--- 11. sync.Once ---")

	var once sync.Once
	var wg sync.WaitGroup

	initialization := func() {
		fmt.Println("Initializing (this should only happen once)")
	}

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			once.Do(initialization)
		}()
	}

	wg.Wait()
}

// 12. Pipeline Pattern
func demoPipelinePattern() {
	fmt.Println("\n--- 12. Pipeline Pattern ---")

	// Stage 1: Generate numbers
	gen := func() <-chan int {
		out := make(chan int)
		go func() {
			for i := 0; i < 10; i++ {
				out <- i
			}
			close(out)
		}()
		return out
	}

	// Stage 2: Square numbers
	square := func(in <-chan int) <-chan int {
		out := make(chan int)
		go func() {
			for n := range in {
				out <- n * n
			}
			close(out)
		}()
		return out
	}

	// Stage 3: Print results
	print := func(in <-chan int) {
		for n := range in {
			fmt.Println("Pipeline result:", n)
		}
	}

	// Connect the pipeline
	print(square(gen()))
}

// 13. Pub/Sub Pattern
func demoPubSubPattern() {
	fmt.Println("\n--- 13. Pub/Sub Pattern ---")

	type Message struct {
		Topic   string
		Content string
	}

	broker := make(chan Message)
	done := make(chan bool)

	// Subscriber
	go func() {
		subscriptions := make(map[string]chan Message)
		for {
			select {
			case msg := <-broker:
				if ch, exists := subscriptions[msg.Topic]; exists {
					ch <- msg
				}
			case <-done:
				return
			}
		}
	}()

	// Publisher
	publish := func(topic, content string) {
		broker <- Message{Topic: topic, Content: content}
	}

	// Subscribe function
	subscribe := func(topic string) chan Message {
		ch := make(chan Message)
		go func() {
			broker <- Message{Topic: topic, Content: "SUBSCRIBE"}
			// This is simplified - real implementation would need proper registration
		}()
		return ch
	}

	// Example usage
	weatherChan := subscribe("weather")
	go func() {
		for msg := range weatherChan {
			fmt.Printf("Weather update: %s\n", msg.Content)
		}
	}()

	publish("weather", "Sunny with a chance of goroutines")
	publish("weather", "Heavy concurrency expected")

	time.Sleep(100 * time.Millisecond)
	done <- true
}
