package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int8
}

func main() {
	p1 := Person{name: "defulte", age: 0}

	p1.age = 24
	p1.name = "parth"

	fmt.Println(p1)
	udpateProfile(&p1, "dgd")
}

func udpateProfile(s *Person, updateName string) {
	fmt.Println(s.name, (*s).name)
	s.name = updateName
}
