package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Please rate our app")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	// first trim input after that convert into float value 
	number, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		incre := number + 1.5
		fmt.Println(incre)
	}
}
