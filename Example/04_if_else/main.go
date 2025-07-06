package main

import "fmt"

func main() {
	var age int
	fmt.Print("Please Enter your age: ")
	_, err := fmt.Scanf("%d", &age)

	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	if age < 18 {
		fmt.Println("You are too young")
	} else if age == 18 {
		fmt.Println("you are perfect age")
	} else {
		fmt.Println("You are older than you have too")
	}
}
