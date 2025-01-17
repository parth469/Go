package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("Hello Word")
	var s string = "str"
	li := strings.Split(s, "")
	fmt.Println(strings.Join(li, "->"))
	var age int = 234
	var price float64 = 32.3523
	var isActive bool = true
	list := []string{"234", "234234", "234"}
	const a = "234d"
	f := "str"
	num, err := strconv.Atoi(s)
	if err != nil {
		fmt.Errorf("Fail to Convet Number")
	}
	fmt.Println(num)
	parth := Person{Name: "parth", Age: 234}
	fmt.Println(s, age, price, isActive, f, list, parth.Age, parth)

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	start := 2
	end := 15
	for start < end {
		fmt.Println(start)
		start++
		end--
	}

	for index, val := range list {
		fmt.Println(index, val)
		if val == "234234" {
			break
		}
	}

	ff := false && true || false
	fmt.Println(ff)
	i := 23
	switch i {
	case 2:
		fmt.Println("It 2")
	case 25:
		fmt.Println(25)
	default:
		fmt.Println(i)
	}

	mp := make(map[string]string)

	mp["new"] = "234"
	delete(mp, "new")
	for key, value := range mp {
		fmt.Println(key, value)
	}
	vari("34", 234, 234, 234, 23, 4)
}

func add(a, b int) int {
	return a + b
}

func sub(num1, num2 int) (sum int) {
	sum = num1 - num2
	return
}

func mit(str string, num int) (st string, nu int) {
	st = str
	nu = num
	return
}

func vari(str string, nums ...int) {
	sum := 0
	for _, s := range nums {
		sum += s
	}
	fmt.Println(str, sum)
}
