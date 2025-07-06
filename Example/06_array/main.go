package main

import (
	"fmt"
	"slices"  // Go 1.21+ required for slices.Reverse and other slice utils
	"sort"
	"strings"
)

func main() {
	// Initialize a string slice (array in Go is fixed-size, slice is more flexible)
	arr := []string{"Apple", "banana", "strabarry"}

	// Sort 'arr' in descending order based on string length
	sort.Slice(arr, func(i, j int) bool {
		return len(arr[i]) > len(arr[j]) // longer strings come first
	})

	// Print joined string with "-"
	news := strings.Join(arr, "-")
	fmt.Println("Joined string:", news) // e.g., "strabarry-banana-Apple"

	// Print the last element and the total length
	fmt.Println("Last element:", arr[len(arr)-1], "Length:", len(arr))

	// Append new elements using spread operator (...) for variadic expansion
	arr = append(arr, []string{"test", "he"}...)
	fmt.Println("After append:", arr)

	// Create and sort a slice of integers
	arrs := []int{1, 3, 5, 6}

	// ❗ This logic is incorrect: it compares index instead of values
	// The correct way is: arrs[i] > arrs[j]
	sort.Slice(arrs, func(i, j int) bool {
		return arrs[i] > arrs[j] // sort in descending order
	})

	// Make a copy of the sorted slice
	arrsCopy := make([]int, len(arrs))
	copy(arrsCopy, arrs)

	// Print index and value of arrs
	for i, v := range arrs {
		fmt.Println("Index:", i, "Value:", v)
	}

	// Make another copy
	copyArray := make([]int, len(arrs))
	copy(copyArray, arrs)

	// Reverse the original slice
	slices.Reverse(arrs)

	// Print reversed and original copies
	fmt.Println("Reversed arrs:", arrs)
	fmt.Println("Unmodified copy:", copyArray)
}
