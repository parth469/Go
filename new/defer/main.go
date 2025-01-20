package main

import "fmt"

func sum(num1 int, num2 int) (Total int) {
	Total = num1 + num2
	defer fmt.Println(Total, "sum")
	return
}

func product(num1 int, num2 int) (prod int) {
	prod = num1 * num2
	defer fmt.Println(prod, "pro")
	return
}

func main() {
	defer fmt.Println("Defer call 1")
	defer fmt.Println("Defer call2")
	defer fmt.Println("Defer call 3")
	suma := sum(1, 4)
	defer fmt.Println(suma)
	prod := -1
	defer func() {
		prod = product(1, 4)
		fmt.Println(prod)
	}()
	fmt.Println(prod)
	for i := 0; i < 5; i++ {
		defer fmt.Println(i, "hi")
	}
}
