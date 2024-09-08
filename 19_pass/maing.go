package main

import "fmt"

// In Go, everything is passed by value.
// However, for slices and maps, the value passed is a pointer to the underlying data structure,
// so modifying them in a function can affect the original.

func main() {
	str := "hello"
	printValue(str)
	fmt.Println("After printValue, str:", str) // str remains unchanged

	mp := make(map[string]string)
	mp["TS"] = "Typescript"
	updateMap(mp)
	fmt.Println("After updateMap, mp:", mp) // mp is updated with new entry
	slice := []int{1, 2, 3}
	updateSlice(slice)
	fmt.Println("After updateSlice, slice:", slice) // Original slice is affected for modifications within the length

	appendToSlice(slice)
	fmt.Println("After appendToSlice, slice:", slice) // Original slice is not affected by appends exceeding capacity

	removeElement(slice, 1)

	fmt.Println("After removeElement, slice:", slice)
}

// pass by value (for basic types like strings)
func printValue(str string) {
	str = "no" // This change is local and won't affect the original 'str'
	fmt.Println("Inside printValue, str:", str)
}

// since maps are reference types, the changes made here reflect on the original map
func updateMap(mp map[string]string) {
	mp["JS"] = "JavaScript" // Modifies the original map
	fmt.Println("Inside updateMap, mp:", mp)
}

// Modify elements within the original slice length
func updateSlice(slice []int) {
	slice[0] = 10 // This modification affects the original slice
	fmt.Println("Inside updateSlice, slice:", slice)
}

// Appending beyond the slice's capacity
func appendToSlice(slice []int) {
	slice = append(slice, 4, 5, 6) // This creates a new underlying array
	slice[0] = 100                 // Changes to this new slice do not affect the original
	fmt.Println("Inside appendToSlice, slice:", slice)
}

func removeElement(slice []int, index int) []int {
	slice = append(slice[:index], slice[index+1:]...)
	fmt.Println(slice)
	return slice
}
