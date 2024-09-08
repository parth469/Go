package main

import "fmt"

func main() {
	fmt.Println("Welcome to type")
	task := TodoList{"Hi", "hi"}
	fmt.Println(task)
	parth := Person{Name: "parth", Age: 34}
	detail := parth.GetPersonDetail()
	fmt.Println(detail)
	c1 := Circle{Redis: 24}
	fmt.Println(c1.getArea())
}

func (p Person) GetPersonDetail() string {
	return fmt.Sprint("Name ", p.Name, " Age ", p.Age)
}

type TodoList []string

type Person struct {
	Name string
	Age  int
}

type Circle struct {
	Redis float64
}

const PI = 3.14
// by defult is copy of data 
// to update original use pointer 
func (c Circle) getArea() float64 {
	return PI * c.Redis * c.Redis
}
// Method Overload not possbile 
// func (c1, c3 Circle) getArea() float64 {
// 	return PI * c1.Redis * c1.Redis
// }