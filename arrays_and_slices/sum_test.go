package arrays_and_slices_test

import (
	"learn-go-with-tests/arrays_and_slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := arrays_and_slices.Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d when given %v", got, want, numbers)
		}
	})
}
