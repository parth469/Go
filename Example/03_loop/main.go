package main

import "fmt"

func main() {

	// Basic Loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// while loop
	i := 0
	for i < 3 {
		i++
		fmt.Println(i, "while")
	}

	// range
	s := "hi there"

	for i, v := range s {
		fmt.Println(i, string(v))
	}

	mp := map[string]string{"a": "apple", "b": "banana"}

	for k, v := range mp {
		fmt.Println("key is", k, "and value is ", v)
	}

	// array interval
	b := [3]string{"apple", "banana", "cherry"}
	for i, v := range b {
		fmt.Println("index", i, "value is ", v)
	}
}
