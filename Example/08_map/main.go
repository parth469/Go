package main

import (
	"fmt"
	"maps"
	"slices"
)

func main() {
	// Initializing maps in different ways
	mp := make(map[string]int) // Using make()
	m := map[string]int{}      // Using literal syntax

	// Creating a clone of an empty map (this is safe even if m is nil)
	cloneMap := maps.Clone(m)

	// Basic map operations
	mp["k1"] = 7       // Add key-value pair
	mp["k1"] = 747     // Update existing key
	cloneMap["k1"] = 6 // This modifies the clone, not the original
	mp["k2"] = 13
	mp["k3"] = 0

	// DeleteFunc removes entries where the function returns true
	// This example would remove odd values (but m is empty in this case)
	maps.DeleteFunc(m, func(k string, v int) bool {
		return v%2 != 0
	})

	// Getting all keys from a map (order is random)
	fmt.Println("All keys in mp:", maps.Keys(mp))
	fmt.Println("Current mp:", mp)

	// Getting and sorting values
	val := maps.Values(m)        // Get all values
	values := slices.Sorted(val) // Sort the values
	fmt.Println("Sorted values from m:", values)

	// Insert all key-value pairs from mp into m
	maps.Insert(m, maps.All(mp))
	fmt.Println("After inserting from mp, m is now:", m)

	// More map operations
	fmt.Println("Current mp:", mp, "Length:", len(mp))

	// Iterating through a map (order is random)
	for k, v := range mp {
		fmt.Printf("Key: %s, Value: %d\n", k, v)
	}

	// Accessing values
	fmt.Println("k1 value:", mp["k1"])

	// Safe value access with existence check
	value, exists := mp["k3"]
	fmt.Printf("Value: %d, Exists: %t\n", value, exists)

	// Deleting a key
	delete(mp, "k3")
	fmt.Println("After deleting k3:", mp)

	// Clearing a map by deleting all keys
	for k := range mp {
		delete(mp, k)
	}
	fmt.Println("After clearing, mp:", mp, "Clone:", cloneMap)

	// Additional useful operations you might want to consider:

	// 1. Check if two maps are equal (have the same key-value pairs)
	fmt.Println("Are m and mp equal?", maps.Equal(m, mp))

	// 2. Copy from source to destination (like Insert but overwrites existing keys)
	maps.Copy(m, cloneMap)
	fmt.Println("After copying from cloneMap to m:", m)

	// 3. Check if a map is empty
	fmt.Println("Is mp empty?", len(mp) == 0)

	// 4. Create a map with initial capacity (optimization for large maps)
	largeMap := make(map[string]int, 1000)

	// 5. Using maps with struct keys
	type Coord struct{ x, y int }
	coordMap := make(map[Coord]string)
	coordMap[Coord{1, 2}] = "point A"
	fmt.Println("Coord map:", coordMap)

	// 6. Nested maps
	nested := make(map[string]map[string]int)
	nested["outer"] = make(map[string]int)
	nested["outer"]["inner"] = 42
	fmt.Println("Nested map:", nested, largeMap)
}
