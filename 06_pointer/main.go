package main

import "fmt"

func main() {
	// String variable initialized with value "Hi"
	pass := "Hi"

	// Passing the address of the `pass` variable to `passbyValue` function
	// This will modify the original string since it's passed by reference (pointer)
	passbyValue(&pass)
	fmt.Println(pass) // Prints: slfdj (modified by the function)

	// Calling the variadic function with multiple integer arguments
	totalSum := variadicFun(4, 5, 3)
	fmt.Println(totalSum) // Prints the sum: 12

	// Creating and initializing a map
	langMp := make(map[string]string)
	langMp["JS"] = "javascript"
	langMp["TS"] = "Typescript"

	// Printing the map before modification
	fmt.Println(langMp) // Prints: map[JS:javascript TS:Typescript]

	// Passing the map to the `updateMap` function
	// Maps are reference types, so changes inside the function affect the original map
	updateMap(langMp)

	// After deletion in `updateMap`, the original map in `main` is modified
	fmt.Println(langMp) // Prints: map[JS:javascript], as "TS" was deleted

	// Slice Example: Slices are also reference types like maps
	sliceExample := []int{1, 2, 3, 4, 5}

	// Passing the slice to a function that modifies it
	updateSlice(sliceExample)

	// The original slice is modified because slices are passed by reference
	fmt.Println(sliceExample) // Prints: [1 2 100 4 5], as the third element was modified

}

// Function that takes a pointer to a string and modifies the original value
func passbyValue(str *string) {
	*str = "slfdj"    // Modifying the value pointed to by the string pointer
	fmt.Println(*str) // Prints: slfdj
}

// Function that takes a map and modifies it (by deleting a key-value pair)
// Maps are reference types, so this affects the original map in the calling function
func updateMap(maps map[string]string) {
	delete(maps, "TS") // Deletes the "TS" key from the map
	fmt.Println(maps)  // Prints the modified map: map[JS:javascript]
}



// Variadic function that sums integers passed as arguments
// The arguments are passed as a slice (i.e., the `as` variable is a slice of integers)
func variadicFun(as ...int) (sum int) {
	for _, val := range as {
		sum += val // Summing each value
	}
	return // Returns the total sum
}

// Function that modifies a slice by changing its third element
// Slices, like maps, are reference types, so this affects the original slice
func updateSlice(slice []int) {
	slice[2] = 100     // Modifying the third element of the slice
	fmt.Println(slice) // Prints the modified slice: [1 2 100 4 5]
}
