package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	slice := []int{1, 2, 4, 5, 6}

	// Creating a destination slice with the same length as the source slice
	dst := make([]int, len(slice))

	// Append a value to the beginning of the slice
	jds := append([]int{35}, slice...)

	// Copy elements from the source slice to the destination slice
	copy(dst, slice)

	// Function to modify the first element of the slice
	test(dst)

	// Print each element of the slice using a traditional for loop
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	// Print index and value using a range loop
	for i, v := range slice {
		fmt.Println(i, v)
	}

	// Compare two slices
	compareResult := slices.Compare(dst, slice)
	fmt.Println("Comparison result:", compareResult)

	// Find the index of an element in the slice (or -1 if not found)
	fmt.Println("Index of 4 in slice:", slices.Index(slice, 4))

	// Manually find the index of an element (alternative way)
	index := -1
	for i := 0; i < len(slice); i++ {
		if slice[i] == 4 {
			index = i
		}
	}
	fmt.Println("Manual index of 4:", index)

	// Use IndexFunc to find an element based on a condition
	fmt.Println("Index of element equal to 44:", slices.IndexFunc(slice, func(n int) bool {
		return n == 44
	}))

	// Find and delete an element in the slice
	indexDelete := slices.Index(slice, 4)
	deleteSlice := slices.Concat(slice[:indexDelete], slice[indexDelete+1:]) // Concatenate slices excluding the element to be deleted

	// Print original, destination, and modified slices
	fmt.Println("Original slice:", slice)
	fmt.Println("Destination slice after copy:", dst)
	fmt.Println("Slice after deletion of 4:", deleteSlice)
	fmt.Println("Slice with 35 prepended:", jds)

	// Working with a slice of maps
	sli := make([]map[string]int, 1)
	mp1 := map[string]int{"234": 3}
	sli = append(sli, mp1)

	// Sort the slice 'jds' in descending order using SortFunc
	fmt.Println("Slice before sorting:", jds)
	slices.SortFunc(jds, func(last, current int) int {
		if last <= current {
			return 1
		} else {
			return -1
		}
	})

	stringSlice := []string{}

	for _, v := range dst {
		stringSlice = append(stringSlice, string(v))
	}

	fmt.Println(stringSlice)

	for _, v := range stringSlice {
		fmt.Println(v)
	}

	fmt.Printf("%T type of list", stringSlice)

	sort.Slice(dst, func(i, j int) bool {
		return i > j
	})
	fmt.Println(dst, "hello")

	// Print the sorted slice
	fmt.Println("Sorted slice:", jds)
}

func test(arr []int) {
	// Modify the first element of the slice (demonstrates slices are reference types)
	arr[0] = 99
}
