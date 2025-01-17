package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome to struck")
	var foo Person
	fmt.Println(foo) // zero value struct
	parth := Person{Name: "parth", status: true, Age: 21, DateofBirth: time.Date(2001, time.August, 21, 0, 0, 0, 0, time.UTC), class: map[string]string{"ac": "sdf"}}
	fmt.Println("parth detail ", parth)
}

type Person struct {
	Name        string
	Age         int
	status      bool
	DateofBirth time.Time
	class       map[string]string
}
