package main

import (
	"fmt"
)

func main() {
	fmt.Println("Learning Map")

	// Example of a nil map (uninitialized map).
	// Declaring a nil map doesn't allocate memory, so assigning values to it
	// will cause a runtime error.
	var nilVal map[string]int
	// Uncommenting the below line will cause a runtime panic because `nilVal` is not initialized.
	// nilVal["sddf"] = 234
	fmt.Println("Nil map:", nilVal)

	// Example of initializing an empty map using the map literal.
	// This is a valid empty map where you can safely add key-value pairs.
	mp := map[string]int{}
	mp["sddf"] = 234
	fmt.Println("Initialized empty map with a value:", mp)

	// Example of using `make` to initialize a map.
	// `make` allocates memory and initializes the map, making it safe to add values.
	lang := make(map[string]string)
	lang["JS"] = "JavaScript"
	lang["TS"] = "TypeScript"
	lang["py"] = "Python"
	fmt.Println("Map with programming languages:", lang)

	// Example of checking if a key exists in the map.
	// The second return value indicates whether the key was found or not.
	langFind, found := lang["JS"]
	if found {
		fmt.Println("Found language for key 'JS':", langFind)
	} else {
		fmt.Println("Key 'JS' not found")
	}

	// Deleting a key from the map.
	delete(lang, "TS")
	fmt.Println("Map after deleting key 'TS':", lang)

	// Iterating over a map using a `for` loop.
	for key, value := range lang {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}

	// Example of iterating over a string (a string is essentially a slice of runes).
	// It prints the index and corresponding rune (character) in the string.
	const k string = "GoLang"
	for index, value := range k {
		fmt.Printf("Index: %d, Rune: %c, Char: %s\n", index, value, string(value))
	}
}
