package main

import (
	"fmt"
	"math"
)

func variadci(numbers ...int) int {
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	return sum
}

func Closure(multiPier int) func(int) int {
	value := 0

	fn := func(number int) int {
		value = multiPier * number
		return value
	}

	return fn
}

func factorial(number int) int {

	if number == 0 {
		return 1
	}

	return factorial(number-1) * number
}

func deferCallFUnction() {
	fmt.Print("Other Defer")
}

func deferExample() {
	defer fmt.Println("Main defer function")
	fmt.Println("Inside main Defer")
	defer deferCallFUnction()
	defer fmt.Println("end of defer")
}

func square(x *float64) {
	*x = *x * *x
}

func swap(x, y *int) {
	*x, *y = *y, *x
}

type Circle struct {
	x, y, r float64
}

func (c Circle) areasdfa() float64 {
	return math.Pi * c.r * c.r
}

type Person struct {
	Name string
}

type Android struct {
	Person
	Model string
}

func (p Person) Talk() {
	fmt.Println("Hi my name is", p.Name)
}

type Cycles struct {
	r float64
}

type Ractengal struct {
	w float64
	h float64
}

type Square struct {
	w float64
}

func (c Cycles) area() float64 {
	return math.Pi * c.r * c.r
}

func (r Ractengal) area() float64 {
	return r.w * r.h
}

type Shap interface {
	area() float64
}

func mainS() {
	sum1 := variadci(1, 2, 3, 4, 5)

	sl := []int{1, 2, 3, 4, 5}
	sum2 := variadci(sl...)

	fmt.Println(sum1, sum2)

	double := Closure(2)
	twintiyTime := Closure(20)
	result := double(4)

	fmt.Println(result, twintiyTime(3))

	fmt.Println(factorial(5))

	deferExample()
	x := 1.5
	square(&x)

	fmt.Println(x)

	s, y := 1, 2
	swap(&s, &y)

	fmt.Println(s, y)

	c := Circle{y: 8, x: 7, r: 8}
	area := c.areasdfa()
	fmt.Println(area)

	android := Android{Person: Person{Name: "Parth"}, Model: "1"}
	android.Talk()
}
