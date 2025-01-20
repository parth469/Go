package main

import (
	"encoding/json"
	"fmt"
)

type myData struct {
	Name string `json:"full_name"`
	Age  int `json:"person_age"`
}

func mains() {
	parth := myData{Name: "parth", Age: 1}

	empJSON, err := json.Marshal(parth)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(empJSON))
	// khushi := &parth
	// parth.age = 234
	// fmt.Println(khushi.age, parth)

	// em := new(myData)
	// // ks := *em  only copy value
	// ks := em // copy pointer
	// ks.age = 234
	// fmt.Println(ks, em)
	// fmt.Printf("Emp Pointer: %p\n", em)
	// fmt.Printf("%+v", *em) // pritn in {name: age:234} formate
	// fmt.Print("%v", *em)

	// // khushi := myData{"khushi"} raise error have to provide all the value without key
	// payal := myData{23, "khushi"} error with key order of value have to same order
	// mayuri := myData{}
	// fmt.Println(parth, mayuri)
	// mayuri.age = 23
	// fmt.Println(mayuri)

}
