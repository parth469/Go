package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type employee struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	jsonData := `{"someXYZ":"John","age":21}`

	var emp employee

	// Use a json.Decoder to parse the input
	decoder := json.NewDecoder(strings.NewReader(jsonData))
	// Enable strict unmarshalling
	decoder.DisallowUnknownFields()

	// Attempt to decode the JSON into the struct
	err := decoder.Decode(&emp)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Decoded struct:", emp)
	}
}
