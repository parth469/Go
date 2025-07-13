package main

import (
	"fmt"
)

type Human struct {
	isAlive bool
}
type Person struct {
	name string
	age  int8
	Human
}



func (p *Person) IsPersionAlive() bool {
	return p.isAlive
}

func main() {
	p1 := Person{name: "defulte", age: 0, Human: Human{isAlive: false}}
	fmt.Println(p1.Human.isAlive)
	p1.age = 24
	p1.name = "parth"

	fmt.Println(p1)
	udpateProfile(&p1, "dgd")
}

func udpateProfile(s *Person, updateName string) {
	fmt.Println(s.name, (*s).name)
	s.name = updateName
}
