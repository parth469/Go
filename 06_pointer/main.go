package main

import "fmt"

func main() {
	pass := "Hi"
	passbyValue(&pass)
	fmt.Println(pass)
}

func passbyValue(str *string) {
	*str = "slfdj"
	fmt.Println(*str)
}
