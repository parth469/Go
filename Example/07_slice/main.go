package main

import (
	"fmt"
	"slices"
)

func main() {
	// Create a slice of integers with length 6, all initialized to 0
	slice := make([]int, 6)
	fmt.Println("Initial int slice (length 6):", slice)

	// Appending an element — increases length to 7
	slice = append(slice, 1)
	fmt.Println("After append 1:", slice)

	// Create a second slice of integers with length 5, default values = 0
	sliceD := make([]int, 5)
	fmt.Println("sliceD:", sliceD)

	// Create a slice of booleans, default values = false
	sliceBool := make([]bool, 5)
	fmt.Println("sliceBool:", sliceBool)

	// Create a slice of strings, default values = ""
	sliceStr := make([]string, 5)
	fmt.Println("sliceStr:", sliceStr)

	// Slice the original `slice` from index 1 to 4 (5 is exclusive)
	l := slice[1:5] // slice from index 1 up to but not including 5
	fmt.Println("Sub-slice l (slice[1:5]):", l)

	// Modify index 3 of the original slice
	slice[3] = 23
	fmt.Println("After slice[3] = 23:", slice)

	// Access safe indexes for printing
	fmt.Println("Accessing individual elements safely:")
	fmt.Println("sliceD[2]:", sliceD[2])
	fmt.Println("sliceStr[0]:", sliceStr[0])
	fmt.Println("sliceBool[0]:", sliceBool[0])
	sl := []int{1, 2, 3, 4, 5, 6, 7}
	app(sl[1:3])
	fmt.Println(sl)
}
func app(sl []int) {
	sl = append(sl, 0)
	slices.Delete(sl, 1, 2)
	fmt.Println(sl)
}
