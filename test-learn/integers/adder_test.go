package integers

import "testing"

func TestAddReturnCorrectValue(t *testing.T) {
	number1, number2 := 1, 5
	got := Add(number1, number2)
	want := number1 + number2

	if got != want {
		t.Errorf("Add(%d, %d) = %d; want %d", number1, number2, got, want)
	}
}

func TestAddReturnCorrectValueWithNagative(t *testing.T) {
	number1, number2 := 1, -5
	got := Add(number1, number2)
	want := number1 + number2

	if got != want {
		t.Errorf("Add(%d, %d) = %d; want %d", number1, number2, got, want)
	}
}

func TestAddReturnCorrectWrongValue(t *testing.T) {
	number1, number2 := 1, 5
	got := Add(number1, number2)
	want := number1 + number2 + 1

	if got == want {
		t.Errorf("Add(%d, %d) = %d; want %d", number1, number2, got, want)
	}
}
