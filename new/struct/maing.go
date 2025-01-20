package main

import "fmt"

func main() {
	p1 := Person{age: 12, Name: "Nage"}
	p2 := &p1
	p1.getDetail("new Name")
	fmt.Println(p1, p2)
}

type Person struct {
	Name string
	age  int
}

func (P *Person) getDetail(s string) {
	P.Name = s
}
