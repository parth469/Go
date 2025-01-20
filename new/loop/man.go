package main

import "fmt"

func main() {
	letters := "casd"
	for index, value := range letters {
		fmt.Println(index, string(value))
	}

	mp := map[string]int{
		"a": 1,
		"b": 3,
	}

	for index, value := range mp {
		fmt.Println(index, value)
	}

	a := 5
	switch a {
	case 1:
		fmt.Println("it 1")
	case 2:
		fmt.Println("two")
	case 3, 5:
		fmt.Println("three or five")
	default:
		fmt.Println("defalult")
	}
}

func test3() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

func test1() {
	for i := 0; i < 5; {
		i = i + 1
		fmt.Println(i)
	}
}
func test2() {
	for i := 0; ; {
		if i > 5 {
			break
		}
		i = i + 1
		fmt.Println(i)
	}
}
func test() {
	i := 0
	for {
		if i > 5 {
			break
		}
		i = i + 1
		fmt.Println(i)
	}
}
