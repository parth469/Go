// Mistake 01: Unintended variable shadowing.
// Re-declaring a variable with := in an inner scope creates a new variable,
// which can hide the outer one and cause unexpected behavior.

package commonmistake

import "fmt"

func problem01() {
	personName := "parth"
	func() {
		// BUG: This creates a new local variable that shadows personName.
		personName := "suraj"
		fmt.Println("[problem] inside block:", personName)
	}()

	// The outer variable is unchanged because the inner one was shadowed.
	fmt.Println("[problem] outside block:", personName)
}

func fix01() {
	personName := "parth"
	func() {
		// FIX: Reuse the outer variable with assignment (=), not re-declaration (:=).
		personName = "suraj"
		fmt.Println("[fix] inside block:", personName)
	}()

	fmt.Println("[fix] outside block:", personName)
}

func Mistake01() {
	fmt.Println("Mistake 01: Variable shadowing in inner scope")
	problem01()
	fix01()
}
