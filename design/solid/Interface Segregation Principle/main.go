package main

import "fmt"

// ✅ Correct: Separate interfaces
type Workable interface {
	Work()
}

type Eatable interface {
	Eat()
}

// ✅ Developer implements both Workable and Eatable
type Developer struct{}

func (d Developer) Work() {
	fmt.Println("Developer is coding...")
}

func (d Developer) Eat() {
	fmt.Println("Developer is eating...")
}

// ✅ Robot only implements Workable, not Eatable
type Robot struct{}

func (r Robot) Work() {
	fmt.Println("Robot is working...")
}

func WorkTask(w Workable) {
	w.Work()
}

func MealBreak(e Eatable) {
	e.Eat()
}

func main() {
	dev := Developer{}
	robot := Robot{}

	WorkTask(dev)  // ✅ Works fine
	MealBreak(dev) // ✅ Works fine

	WorkTask(robot)  // ✅ Works fine
	// MealBreak(robot) // ❌ Compile-time error (Robot does not eat)
}
