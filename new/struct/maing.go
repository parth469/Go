package main

import "fmt"

func maidns() {

	type Person struct {
		Name string
		age  int
	}

	p1 := &Person{"parth", 23}
	p2 := p1
	fmt.Printf("%p type is and Person Type is %p", p1, p2)
}
