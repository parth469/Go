package main

import "fmt"

// Variadic function `sum` accepts any number of int arguments and returns their total sum.
func sum(arg ...int) int {
	result := 0
	for _, v := range arg {
		result += v
	}
	return result
}

// `multiply` takes two integers and returns their product.
func multiply(n1, n2 int) int {
	return n1 * n2
}

// `multiReturn` demonstrates multiple return values.
// It takes a name and age, returns a message string and modified age.
func multiReturn(name string, age int) (string, int) {
	updateName := "your name is " + name
	return updateName, age + 19
}

// `closerExample` returns a closure.
// A closure is a function that "remembers" the variables in its lexical scope.
func closerEXample() func() int {
	i := 0
	// The returned function increments and returns `i` every time it's called.
	return func() int {
		i++
		return i
	}
}

// `Recursion` calculates factorial using recursion.
// Factorial(n) = n * (n-1) * (n-2) ... * 1
func Recursion(i int) int {
	if i == 1 {
		return 1
	}
	return i * Recursion(i-1) // Recursive case
}

func main() {
	// Call to variadic sum function
	totalSum := sum(1, 5, 6, 7)

	// Calling multiReturn normally
	updateName, updateAge := multiReturn("parth", 14)

	// Calling multiReturn using an anonymous function (inline declaration)
	first, second := func() (string, int) {
		return multiReturn("parth", 14)
	}()

	fmt.Println(first, second)          // Output from anonymous function call
	fmt.Println(updateAge, updateName)  // Output from normal call to multiReturn
	fmt.Println(totalSum)               // Total sum of numbers passed to `sum`

	// Closure example
	fun := closerEXample()
	fmt.Println(fun()) // 1
	fmt.Println(fun()) // 2
	fmt.Println(fun()) // 3
	fmt.Println(fun()) // 4
	fmt.Println(fun()) // 5

	// New closure instance - separate state from `fun`
	Secfun := closerEXample()
	fmt.Println(Secfun(), fun()) // Secfun() = 1, fun() = 6 (continues from previous state)

	// Recursive function - factorial of 5 => 5*4*3*2*1 = 120
	fmt.Println(Recursion(5)) // Output: 120
}
