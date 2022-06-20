package ch4_test

import (
	"learn-go-with-tests/ch4"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := ch4.Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d when given %v", got, want, numbers)
		}
	})
}

func checkSums(t testing.TB, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAll(t *testing.T) {
	t.Run("2 slices", func(t *testing.T) {
		got := ch4.SumAll([]int{1, 2}, []int{3, 4})
		want := []int{3, 7}

		checkSums(t, got, want)
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("Slices with values", func(t *testing.T) {
		got := ch4.SumAllTails([]int{1, 2, 3}, []int{4, 5, 6})
		want := []int{5, 11}

		checkSums(t, got, want)
	})

	t.Run("Empty slices", func(t *testing.T) {
		got := ch4.SumAllTails([]int{}, []int{})
		want := []int{0, 0}

		checkSums(t, got, want)
	})
}
