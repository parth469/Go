package main

import "fmt"

// 1. Modify: Changes reflect in original slice
func modifySlice(s []int) {
	s[0] = 99 // Modifies original because s shares backing array
}

// 2. Reslice: Changes do NOT reflect unless returned
func reslice(s []int) {
	s = s[1:] // Re-slicing only affects local copy
	fmt.Println("Inside reslice:", s)
}

// 3. Append: Changes do NOT reflect unless returned
func appendSlice(s []int) {
	s = append(s, 100) // Appends to local copy
	fmt.Println("Inside appendSlice:", s)
}

// 4. Append in-place: Only reflects if capacity not exceeded
func appendInPlace(s []int) {
	s[0] = 77           // modify
	s = append(s, 88)   // might or might not reflect outside
	fmt.Println("Inside appendInPlace:", s)
}

func main() {
	// Base slice for all tests
	sl := []int{1, 2, 3}
	sl1 := []int{1, 2, 3}
	sl2 := []int{1, 2, 3}
	sl3 := make([]int, 3, 6) // Capacity is larger to allow in-place append
	copy(sl3, []int{1, 2, 3})

	// 1. Modify
	modifySlice(sl)
	fmt.Println("After modifySlice:", sl) // ✅ [99 2 3]

	// 2. Reslice
	reslice(sl1)
	fmt.Println("After reslice:", sl1)    // ❌ [1 2 3] (unchanged)

	// 3. Append (new array)
	appendSlice(sl2)
	fmt.Println("After appendSlice:", sl2) // ❌ [1 2 3] (unchanged)

	// 4. Append in-place (depends on capacity)
	appendInPlace(sl3)
	fmt.Println("After appendInPlace:", sl3) // May or may not include 88 base on capacity
}
