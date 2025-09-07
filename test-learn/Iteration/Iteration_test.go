package iteration_test

import (
	"testing"

	iteration "github.com/parth469/test/Iteration"
)

func TestIterationReturnCorrectValue(t *testing.T) {
	got := iteration.Iteration("a")
	want := "aaaaa"

	if got != want {
		t.Errorf("Iteration(\"a\") = %s; want %s", got, want)
	}
}

func TestIterationWithNumberAsString(t *testing.T) {
	got := iteration.Iteration("34")
	want := "3434343434"

	if got != want {
		t.Errorf("Iteration(\"34\") = %s; want %s", got, want)
	}
}

func TestIterationWithNilString(t *testing.T) {
	// Go does not allow nil as a string argument, so we test with an empty string instead.
	got := iteration.Iteration("")
	want := ""

	if got != want {
		t.Errorf("Iteration(\"\") = %s; want %s", got, want)
	}
}

func TestIterationWithEmptyString(t *testing.T) {
	got := iteration.Iteration("")
	want := ""

	if got != want {
		t.Errorf("Iteration(\"\") = %s; want %s", got, want)
	}
}

func TestIterationWithMultiCharString(t *testing.T) {
	got := iteration.Iteration("ab")
	want := "ababababab"

	if got != want {
		t.Errorf("Iteration(\"ab\") = %s; want %s", got, want)
	}
}
