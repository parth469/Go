package structMethod

import (
	"math"
	"testing"
)

// TestPerimeter verifies the Perimeter function for various Rectangle dimensions.
func TestPerimeter(t *testing.T) {
	t.Run("Return Correct Result", func(t *testing.T) {
		r := Rectangle{Width: 10.0, Height: 10.0}
		got := Perimeter(r)
		want := 40.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

	t.Run("Negative Width", func(t *testing.T) {
		r := Rectangle{Width: -10.0, Height: 10.0}
		got := Perimeter(r)
		want := 0.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

	t.Run("Negative Height", func(t *testing.T) {
		r := Rectangle{Width: 10.0, Height: -10.0}
		got := Perimeter(r)
		want := 0.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

	t.Run("Negative Width and Height", func(t *testing.T) {
		r := Rectangle{Width: -10.0, Height: -10.0}
		got := Perimeter(r)
		want := 0.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

	t.Run("Zero Height", func(t *testing.T) {
		r := Rectangle{Width: 10.0, Height: 0.0}
		got := Perimeter(r)
		want := 20.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

	t.Run("Zero Width", func(t *testing.T) {
		r := Rectangle{Width: 0.0, Height: 10.0}
		got := Perimeter(r)
		want := 20.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

	t.Run("Zero Width and Height", func(t *testing.T) {
		r := Rectangle{Width: 0.0, Height: 0.0}
		got := Perimeter(r)
		want := 0.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
}

func TestMetod(t *testing.T) {
	t.Run("Rectangle Area Calculation", func(t *testing.T) {
		r := Rectangle{Width: 2.0, Height: 4.0}
		got := r.Area()
		want := 8.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
}

func TestAreaInterface(t *testing.T) {

	checkSum := func(t *testing.T, Shap Shap, result float64) {
		area := Shap.Area()

		if area != result {
			t.Errorf("got %.2f want %.2f", area, result)
		}
	}
	t.Run("Rectangle Area Calculation", func(t *testing.T) {
		r := Rectangle{Width: 2.0, Height: 4.0}
		checkSum(t, r, 8.0)
	})

	t.Run("Cycle Area Calculation", func(t *testing.T) {
		r := Cycle{raids: 4.0}
		checkSum(t, r, 16.0*math.Pi)
	})
}
