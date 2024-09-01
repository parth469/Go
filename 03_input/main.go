package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// var message string = "Welcome to user input , Your name is "
	userInput := bufio.NewReader(os.Stdin)
	fmt.Println("Enter Your Name")

	// err Ok -> we do not have try catch
	input, _ := userInput.ReadString('\n')
	fmt.Printf("type is %T sdfl", input)
}
