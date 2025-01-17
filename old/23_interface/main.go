package main

import "fmt"

func main() {
	c1 := Circle{Radius: 7}
	s1 := Rectangle{Width: 3, Height: 5}

	Shapess := []Shape{c1, s1}
	for _, value := range Shapess {
		fmt.Println("Area", value.Area())
	}
}

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

const PI = 3.14

func (c Circle) Area() float64 {
	return PI * c.Radius * c.Radius
}
