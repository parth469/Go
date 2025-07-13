package main

import "fmt"

type Vehical interface {
	start() string
}

type Car struct {
	name string
}

type Bike struct {
	name string
}

func (c *Car) start() string {
	return fmt.Sprintln(c.name, "start Car")
}

func (b *Bike) start() string {
	return fmt.Sprintln(b.name, "start Bike")
}

func PrintResult(v Vehical) {
	fmt.Println(v.start())
}

func PrintInterface(ve interface{}) {
	switch ve.(type) {
	case int:
		fmt.Println("Your value is Int", ve)
	case bool:
		fmt.Println("Your value is bool", ve)
	case *Car:
		car := ve.(Car)
		car.start()
	case []string:
		fmt.Println("it array of string", ve)
	default:
		fmt.Println("fy")
	}
}
func main() {
	c1 := Car{name: "LABO"}
	b1 := Bike{name: "SUSUKI"}

	inf := []Vehical{&c1, &b1}

	for _, v := range inf {
		PrintResult(v)
	}

	mix := []interface{}{"1", 1, true, []string{"41"}, c1}

	for _, v := range mix {
		PrintInterface(v)
	}

	var empty *int
	// interface to be nil both value and type to be empty
	var emptyI interface{}

	fmt.Println(empty == nil)
	fmt.Println(emptyI == nil)

	emptyI = empty
	fmt.Println(emptyI == nil, emptyI)

}
