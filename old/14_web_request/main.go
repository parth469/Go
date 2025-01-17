package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://reqres.in/api/products/3"

func main() {
	fmt.Println("welcome to go http")
	response, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Printf("Type is %T, ", response)
	// it dev to close by you self
	defer response.Body.Close()
	dataByte, err := io.ReadAll(response.Body)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Printf("type %T \n", dataByte)
	fmt.Println(string(dataByte))
}
