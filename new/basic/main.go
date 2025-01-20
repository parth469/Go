package main

import "fmt"

var global int = 1

func main() {
	// Print a simple message to the console using fmt.Println
	// fmt.Println internally calls fmt.Sprint to convert the arguments into a single string.
	// It then appends a newline character ('\n') and writes the result to the standard output (os.Stdout).
	fmt.Println("Hello World")

	// Declare and initialize variables of various types
	var name string
	name = "parth patel"

	var age int = 365
	var hight, weight int
	hight = 180
	weight = 110

	var hasAnySister, hasAnyBrother bool = false, true

	// Group declaration of variables for better readability
	var (
		hasJob bool = true
		isWFH  bool = true
		salary int  = 50000
	)

	var hasGF = false
	numberOfFriend := 5
	workingHour, isGood := 9, true

	// Using fmt.Println to print multiple values with automatic formatting and newline addition.
	// Each argument is converted to a string using fmt.Sprint, and the results are concatenated with spaces.
	// A newline is appended at the end before the result is written to the standard output.
	fmt.Println(name, "age is", age, "\n height is", hight, "\n and weight is", weight,
		"\n has any sister", hasAnySister, "\n has any brother", hasAnyBrother,
		"\n do you have any gf", hasGF, "\n Number of friends", numberOfFriend)

	// Using fmt.Printf for formatted output.
	// fmt.Printf uses a format string with placeholders to format each argument precisely.
	// The format specifiers (e.g., %s for strings, %d for integers) dictate how each argument is converted to a string.
	// No newline is added automatically; it must be included in the format string if needed.
	fmt.Printf("Name: %s, Age: %d\n", name, age)
	fmt.Printf("Height: %d cm, Weight: %d kg\n", hight, weight)
	fmt.Printf("Has Job: %t, Is WFH: %t, Salary: %d\n", hasJob, isWFH, salary)

	// Using fmt.Print to print without a newline.
	// fmt.Print works similarly to fmt.Println but does not append a newline at the end.
	// Each argument is converted to a string and written to the standard output as-is.
	fmt.Print("Working hour: ", workingHour, ", Feeling good: ", isGood, "\n")

	// Using fmt.Printf again to print the type and value of a variable
	// %T prints the type of the value, and %v prints the value itself.
	// This is useful for debugging or understanding the types of variables in use.
	fmt.Printf("Type of 'numberOfFriend': %T, Value: %v\n", numberOfFriend, numberOfFriend)

	{
		global := 3
		fmt.Println(global) // This will print the local 'global' variable which is 3.
	}
	fmt.Println(global) // This will print the global 'global' variable which is 1.
}
