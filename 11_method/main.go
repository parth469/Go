package main

import (
	"fmt"
	"time"
)

func main() {
	parth := Person{"parth patel", 24, true, "parthnj2001@gmail.com"}
	dateofBirth := parth.GetDateOfBirth()
	fmt.Println("Part Detail", parth, "Date of Birth Year", dateofBirth)
}

type Person struct {
	Name   string
	Age    int
	Status bool
	Email  string
}

func (u *Person) GetDateOfBirth() (birthYear int) {
	currentYear := time.Now().Year()
	u.Name = "parth"
	birthYear = currentYear - u.Age + 1
	return

}
