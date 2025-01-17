package main

import "fmt"

func main() {
	parth := Human{Name: FullName{FirstName: "Parth", LastName: "Patel",MiddleName: "NareshKumar"}, Age: 34, Gender: Male}
	printValue(parth)
}

func (name FullName) getFullName() string {
	return fmt.Sprintf("%s %s %s", name.FirstName, name.MiddleName, name.LastName)
}

func printValue(hu Human)  {
	fmt.Println(hu.Name.getFullName())
}

type Gender string

const (
	Male   Gender = "Male"
	Female Gender = "Female"
	Other  Gender = "Other"
)

// Define the FullName struct
type FullName struct {
	FirstName  string
	LastName   string
	MiddleName string
}

// Define the Human struct with Name and Gender
type Human struct {
	Name   FullName
	Age    int
	Gender Gender
}
