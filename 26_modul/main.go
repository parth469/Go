package main

import (
	"fmt"

	"github.com/fatih/color"
	whymyMath "github.com/parth469/Go/whyMath"
)

func main() {
	colorString := color.RedString("Prints text in cyan.")
	fmt.Println("welcome to module file ", colorString)

	res, error := whymyMath.Min(3, 5)

	if error != nil {
		fmt.Println(error)
	}
	whymyMath.Gooooo()
	fmt.Println(res)
}
