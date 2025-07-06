package main

import "fmt"

func main() {

	name := "parth"
	age := 25

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
	const pi  = 35.53
	// pi = 1415926535  you can not update const value ( it shoude be declure during )

	name = "path has update name"

	fmt.Printf("My name is %s, and age is %d \n", name, age)
	fmt.Printf("%f\n", pi)
	fmt.Printf("%.2f\n", pi)

}
