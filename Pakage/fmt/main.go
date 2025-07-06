package main

import "fmt"

func main() {
	// fmt.Print
	fmt.Println("---> Print")
	fmt.Print("hello", "world\n") //NOTE - Join them without space -> helloworld234

	// fmt.Printf

	name := "parth"
	age := 25

	fmt.Println("---> Printf")

	// %v - default format (any value)
	// %T - type of the value
	// %b - base 2 (binary)
	// %d - base 10 (decimal)
	// %f - decimal floating point
	// %s - basic string
	// %5d - pad with spaces to make 5 chars wide
	// %05d - pad with zeros to make 5 chars wide
	// %.2f - 2 decimal places
	// %5.2f - 5 chars wide with 2 decimal places
	const pi float32 = 3.1415926535

	fmt.Printf("My name is %s, and age is %d \n", name, age)
	fmt.Printf("|%05d|\n", 42) // |00042|
	fmt.Printf("%f\n", pi)
	fmt.Printf("%.2f\n", pi)

	fmt.Println("---> Println")

	fmt.Println("My name is", name, ", and age is", age)

	// Sprint(), Sprintln(), Sprintf() insted print it return string

	s := fmt.Sprint("Hello", "this ", "return ", "string")
	fmt.Printf("%s", s)
}
