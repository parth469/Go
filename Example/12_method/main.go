package main

import "fmt"

type UserStatus string

const (
	UserStatusActive   UserStatus = "Active"
	UserStatusInactive UserStatus = "Inactive"
)

type Person struct {
	Name   string
	Status UserStatus
	Age    int
}

// THis is not inhereis just taking type only just not  method
type Student Person

// Method can have value or pointer receiver
func (p *Person) GetDetail() string {
	p.Name = "name update"
	return fmt.Sprintln("My name is", p.Name, "Age is", p.Age)
}

// Method overloading is not possible
//
//	func (p *Person) GetDetail() int {
//		return p.Age
//	}
func NewPerson(name string, status UserStatus, age int) (Person, error) {
	switch status {
	case UserStatusActive, UserStatusInactive:
		return Person{Name: name, Status: status, Age: age}, nil
	default:
		return Person{}, fmt.Errorf("invalid status: %s", status)
	}
}

func main() {
	p1 := Person{Name: "Parth", Status: "sdfklasdjf", Age: 12} // This will compile, but you want to restrict Status to only allowed values.

	p1, err := NewPerson("Parth", "sdfklasdjf", 12)
	if err != nil {
		panic(err)
	}
	s1 := Student{Name: "Parth", Status: UserStatusActive, Age: 12}
	fmt.Print(s1)
	// sData = s1.GetDetail()  this is not possible

	data := p1.GetDetail()
	fmt.Println(data)
}
