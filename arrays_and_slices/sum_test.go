package arrays_and_slices_test

import (
	"learn-go-with-tests/arrays_and_slices"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}
	got := arrays_and_slices.Sum(numbers)
	want := 15

	if got != want {
		t.Errorf("got %d want %d given %v", got, want, numbers)
	}
}
