package main

import "fmt"

func main() {
	// `cout` is a closure that captures and modifies a counter variable
	cout := counter()

	// `defer` schedules this statement to run when the `main` function is about to return.
	// The `cout(3)` is evaluated immediately, but the `fmt.Println` call is deferred.
	defer fmt.Println(cout(3)) // The result of `cout(3)` (which modifies the counter) is calculated here
	// defer only effect main not inside function 
	
	// Prints the current state of the counter after adding 3 to it
	fmt.Println(cout(3)) // Output: 6 (3 from the deferred call + 3 here)

	// Adds 5 to the current counter value (which is now 6), so it prints 11
	fmt.Println(cout(5)) // Output: 11

	// Adds another 5 to the counter (which is now 11), so it prints 16
	fmt.Println(cout(5)) // Output: 16

	// Inline anonymous function is called immediately, printing "hi"
	func() {
		fmt.Println("hi")
	}()
}

func counter() counterIn {
	counter := 0
	// Returns a closure that captures the `counter` variable and modifies it each time it's called
	return func(in int) int {
		counter += in  // Modifies the captured `counter` variable
		return counter // Returns the modified value
	}
}

type counterIn func(int) int
