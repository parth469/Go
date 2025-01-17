package main

import (
	"fmt"
)

type MultiType interface{}

func main() {
	mixSlice := []MultiType{23, "sdf", 3.35, true}

	for _, val := range mixSlice {
		fmt.Println(val)
	}

	var emptyInterface any
	emptyInterface = "hello"

	str, ok := emptyInterface.(string)

	if ok {
		fmt.Println(str, "Yes")
	}

	// when we check value to multiple then we use type switch
	typeCheck([]string{"sdlfkj"})
}

func typeCheck(value MultiType) {

	switch inputType := value.(type) {
	case int:
		fmt.Println("it", inputType)
	case string:
		fmt.Println("it", inputType)

	case float64:
		fmt.Println("it", inputType)

	case bool:
		fmt.Println("it", inputType)
	default:
		fmt.Println("it", inputType)
	}

}

// type asseertion -> it run time
// string(234) -> type conversion
// emptyInterface.(string) -> type assertion
