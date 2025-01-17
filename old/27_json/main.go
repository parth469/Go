package main

import (
	"encoding/json"
	"fmt"
)

// convet struct to json marshalling rever is unmarshlling

type Role []string
type User struct {
	ID       int    `json:"user_id"`
	Name     string `json:"name,omitempty"`
	Age      int    `json:"age"`
	Password string `json:"-"`
	Role     Role   `json:"roles"`
}

func main() {
	user1 := User{
		ID:       234,
		Name:     "parth Patel",
		Age:      34,
		Password: "23",
		Role:     Role{"3", "34"},
	}

	res, errr := json.Marshal(user1)
	if errr != nil {
		fmt.Println(errr)
	}
	fmt.Println(user1, string(res))
}

// inti function auto run first befor main run
func init() {
	fmt.Println("First")
}
