package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Cannot Divide by zero")
	}
	return float64(a) / float64(b), nil
}

func divides(a, b int) (float64, error) {
	if b == 0 {
		return 0, errors.New("Cannot Divide by zero")
	}
	return float64(a) / float64(b), nil
}

type CustomerError struct {
	Code    int
	Message string
}

// Error implements error.
func (c CustomerError) Error() string {
	return fmt.Sprintf("Error %d, %s", c.Code, c.Message)
}

func divideCs(a, b int) (float64, error) {
	if b == 0 {
		return 0, CustomerError{Code: 1, Message: "Not possbie "}
	}
	return float64(a) / float64(b), nil
}

func main() {
	// res, error := divide(3, 0)
	// if error != nil {
	// 	fmt.Println(error)
	// 	panic(error)
	// }
	// fmt.Println(res)
	error := secondFund()
	fmt.Println(error)
}

func firstFunc() error {

	return fmt.Errorf("original Error")
}

func secondFunc() error {
	err := firstFunc()
	if err != nil {
		return fmt.Errorf("second error %v ", err)
	}
	return nil
}

func secondFund() error {
	err := firstFunc()
	if err != nil {
		return errors.Join(errors.New("this is fail"), err)
	}
	return nil
}
