package commonmistake

import "fmt"

// problem02 shows an unnecessary if-else chain.
// It works, but the nesting is noisier than needed.
func problem02(age int) string {
	if age > 5 && age <= 10 {
		return "too small"
	} else if age > 10 && age <= 18 {
		return "may work"
	} else if age > 18 && age <= 40 {
		return "perfect"
	} else if age > 40 {
		return "too old"
	} else {
		return "wrong"
	}
}

// fix02 keeps the same behavior with flatter, ordered returns.
// This style is easier to scan and maintain.
func fix02(age int) string {
	// Invalid range.
	if age <= 5 {
		return "wrong"
	}

	// From here, age > 5 is guaranteed.
	if age <= 10 {
		return "too small"
	}

	if age <= 18 {
		return "may work"
	}

	if age <= 40 {
		return "perfect"
	}

	return "too old"
}

func Mistake02() {
	fmt.Println("Mistake 02: Multiple if-else / unnecessary nesting")
	age := 20
	problem02(age)
	fix02(age)
}
