package main

import "fmt"

func update(a int) {
	a = a + 10
}

func updateP(a *int) {
	*a = *a + 10
}
func main() {
	var intP *int
	a := 23
	intP = &a
	fmt.Println(intP)
	p := &a
	update(a)
	fmt.Println(a, *p)
	updateP(p)
	fmt.Println(*&a, *p)
}
