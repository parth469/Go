package main

import "fmt"

type Shape interface {
	Area()
}

type Circle struct{}

type Squa struct{}

func (c *Circle) Area() {
	fmt.Println("For Circle")
}

func (c *Squa) Area() {
	fmt.Println("For Squar")
}
func main() {
	ci := Squa{}
	ci.Area()
}
