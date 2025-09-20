package sum

import "testing"

// TestSumReturnsCorrectResult verifies that Sum returns the correct total for a slice of positive integers.
func TestSum(t *testing.T) {
	t.Run("Return Correact Result", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		want := 15
		got := Sum(numbers)

		if got != want {
			t.Errorf("Sum(%v) = %d; want %d", numbers, got, want)
		}
	})

	// TestSumWithNegativeNumbers checks that Sum correctly handles a mix of positive and negative numbers.
	t.Run("Return Correct Result With Negative Number", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, -5}
		want := 5
		got := Sum(numbers)

		if got != want {
			t.Errorf("Sum(%v) = %d; want %d", numbers, got, want)
		}
	})

	// TestSumWithEmptySlice checks that Sum returns 0 for an empty slice.
	t.Run("Return Zero For Empty Slice", func(t *testing.T) {
		numbers := []int{}
		want := 0
		got := Sum(numbers)

		if got != want {
			t.Errorf("Sum(%v) = %d; want %d", numbers, got, want)
		}
	})

	// TestSumWithAllNegativeNumbers checks that Sum correctly sums all negative numbers.
	t.Run("Return Correct Result With All Negative Numbers", func(t *testing.T) {
		numbers := []int{-1, -2, -3, -4, -5}
		want := -15
		got := Sum(numbers)

		if got != want {
			t.Errorf("Sum(%v) = %d; want %d", numbers, got, want)
		}
	})
}
