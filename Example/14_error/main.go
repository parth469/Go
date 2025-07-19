package main

import (
	"errors"
	"fmt"
)

var NotFound error = errors.New("No result Found 0")
var NotFound1 error = errors.New("No result Found 1 ")

// type Error interface {
// 	Error() string
// }

// Error.New
// fmt.ErrorString

func getError(returnValue int) (int, error) {
	if returnValue == 0 {
		return 0, fmt.Errorf("Insie of Erorr")
	}
	return returnValue, nil
}
func main() {
	value, err := getError(0)

	if err != nil {
		fmt.Println(errors.Is(err, NotFound1))
		fmt.Println(errors.Is(err, NotFound))
		// fmt.Println(errors.As(NotFound, NotFound1)) it must be inerface

		errJoin := errors.Join(NotFound, err, NotFound1)
		fmt.Println(errJoin)
		return
	}

	fmt.Println(value)
}
