package iteration

import (
	"strings"
)

var repeatCount = 5

func Iteration(repeatString string) string {
	var repeatStr strings.Builder

	for i := 0; i < repeatCount; i++ {
		repeatStr.WriteString(repeatString)
	}

	return repeatStr.String()
}
