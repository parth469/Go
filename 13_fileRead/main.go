package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to file")
	content := "This is Test file"
	newConetec := "update Content"
	CreateFile(content)
	readFile()
	updateFile(newConetec)
	readFile()

}

func updateFile(newContent string) {
	detail, err := os.OpenFile("./13_fileRead/test.txt", os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer detail.Close()

	// Append new content to the file
	_, err = io.WriteString(detail, newContent)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

}
func readFile() {
	detail, err := os.ReadFile("./13_fileRead/test.txt")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(string(detail))
}

func CreateFile(content string) {
	file, err := os.Create("./13_fileRead/test.txt")
	if err != nil {
		fmt.Println(err, "This error while Creating file")
		panic(err)
	}
	defer file.Close()
	file.Write([]byte(content))
	contentUpdate, err := io.WriteString(file, content)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(contentUpdate)
}
