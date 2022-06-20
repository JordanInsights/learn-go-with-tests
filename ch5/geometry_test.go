package ch5_test

import (
	"learn-go-with-tests/ch5"
	"testing"
)

func assert2fEquality(t *testing.T, got, want float64) {
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestPerimeter(t *testing.T) {
	rectangle := ch5.Rectangle{5.0, 10.0}

	got := ch5.Perimeter(rectangle)
	want := 30.0

	assert2fEquality(t, got, want)
}

func TestArea(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		rectangle := ch5.Rectangle{5.0, 10.0}
		got := rectangle.Area()
		want := 50.0

		assert2fEquality(t, got, want)
	})
	t.Run("circles", func(t *testing.T) {
		circle := ch5.Circle{10}
		got := circle.Area()
		want := 314.1592653589793

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})
}
