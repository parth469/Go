package main

import "fmt"

type Person struct {
	FirstName     string
	Age           int
	LastFirstName string
}

func (e *Person) GetDetailAndStore() (fullName string) {
	fullName = e.FirstName + " " + e.LastFirstName
	fmt.Println(fullName)
	return
}

func main() {
	fmt.Println("Without Single Responsibility")
	parth := Person{LastFirstName: "Patel", FirstName: "Parth", Age: 23}
	parth.GetDetailAndStore()
}
