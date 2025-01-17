package main

import "fmt"

func main() {
	young := Age(34)
	old := Age(74)
	fmt.Println(young + old)

	parth := Human{Age: 34}
	parthSugent := Strudent{Age: 34}
	fmt.Println(parth.Age, parthSugent.Age)
	ex(parth)
	fmt.Println(Monday)
}

func ex(st Human) {
	fmt.Println(st.Age)

}

type Age int

type Human struct {
	Age int
}

type Strudent Human

type Day int

const (
	Sunday    Day = iota // 0
	Monday               // 1
	Tuesday              // 2
	Wednesday            // 3
	Thursday             // 4
	Friday               // 5
	Saturday             // 6
)
