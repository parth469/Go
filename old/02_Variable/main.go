package main

import "fmt"

func main() {
	const bo float32 = 23.23
	str := "code with parth"
	//
	s1 := str[2]  // Access the byte at index 2
	s2 := str[4:] // Slice from index 4 to the end
	s4 := str[:5] // Slice from the start to index 5

	// Print the character, slice from index 4, and slice up to index 5
	fmt.Println(string(s1), s2, s4)

	// Attempting to modify a string like this will cause a compilation error
	// str[3] = "35"

	// If you need to modify a string, create a new one
	newStr := str[:3] + "35" + str[4:]

	fmt.Println(newStr)
	fmt.Println(bo)

	st1 := "हिंदी टाइपिंग टेस्ट"
	fmt.Println(string(st1[1]))

	// ASCII vs. Unicode: ASCII characters are represented by a single byte, but Unicode characters (like those in Hindi) are often represented by multiple bytes (UTF-8 encoding).
	// Convert the string to a slice of runes
	// in unicode it can be 1 byte to 4 byte
	runes := []rune(st1)
	// Access and print the second rune (which is the second character)
	fmt.Println(string(runes[0]))

	newstart := str[:6] + " wjy " + str[6:]
	fmt.Println(newstart)
}
