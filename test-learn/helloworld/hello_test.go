package helloworld

import "testing"

func TestSayHelloReturnsCorrectGreeting(t *testing.T) {
	got := SayHello()
	want := "Hello, world"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestSayHelloDoesNotReturnIncorrectGreeting(t *testing.T) {
	got := SayHello()
	want := "Wrong"

	if got == want {
		t.Errorf("got %q, want %q", got, want)
	}
}
