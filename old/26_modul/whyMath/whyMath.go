package whyMath

import (
	"errors"

	"github.com/fatih/color"
)

func Min(num2, num1 int) (result int, error error) {
	if num2 < 0 {
		errorStr := color.RedString("Please Prvoid Proper Value")
		return -1, errors.New(errorStr)
	}
	result = num1 - num2
	return result, nil
}
