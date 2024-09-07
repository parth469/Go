package main

import (
	"fmt"
)

func main() {
	// Passing the fullName function as an argument to sayHello.
	// fullName will concatenate two strings, and sayHello adds a prefix to it.
	result := sayHello("welcome ", fullName, "parth", "patle")
	fmt.Println(result) // Output: welcome parth patle

	// Using the Access function, which returns a closure.
	// multiplayByTwice is a function that multiplies a number by 2.
	multiplayByTwice := Access(2)

	// multiplayBySiz is a function that multiplies a number by 6.
	multiplayBySiz := Access(6)

	// Calling the returned functions and printing the result.
	fmt.Println(multiplayByTwice(2), multiplayBySiz(6)) // Output: 4 36
}

// fullName takes two strings as input and returns their concatenation.
// It is passed as an argument to the sayHello function.
func fullName(s1, s2 string) (fullname string) {
	fullname = s1 + " " + s2 // Concatenating s1 and s2 with a space in between.
	return
}

// sayHello takes a prefix string, a function, and two additional strings as arguments.
// The function fn is called with the two strings, and its result is concatenated with the prefix.
func sayHello(prefix string, fn func(s1 string, s2 string) (fullname string), s1 string, s2 string) string {
	// Call the function fn with s1 and s2
	full := fn(s1, s2)
	// Return the result of the concatenation of prefix and the full name.
	return prefix + full
}

// Access is a higher-order function that returns a closure (a function).
// The returned function multiplies its input by the mulitplyerNumber passed to Access.
func Access(mulitplyerNumber int) func(int) int {
	// The returned function captures mulitplyerNumber from its surrounding scope (closure).
	return func(num int) int {
		return num * mulitplyerNumber // Multiply the input by the captured multiplier.
	}
}

// Access2 does the same as Access, but uses a type alias for the returned function.
// This improves readability by giving the function type a descriptive name.
func Access2(mulitplyerNumber int) multiplyer {
	// This is identical to Access, but returns a function of type multiplyer (defined below).
	return func(num int) int {
		return num * mulitplyerNumber
	}
}

// multiplyer is a type alias for a function that takes an int and returns an int.
// This makes code more readable when dealing with functions as return types or arguments.
type multiplyer func(int) int
