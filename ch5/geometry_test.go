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
	areaTests := []struct {
		shape ch5.Shape
		want  float64
	}{
		{ch5.Rectangle{12, 6}, 72.0},
		{ch5.Circle{10}, 314.1592653589793},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}

	// checkArea := func(t testing.TB, shape ch5.Shape, want float64) {
	// 	t.Helper()
	// 	got := shape.Area()
	// 	if got != want {
	// 		t.Errorf("got %g want %g", got, want)
	// 	}
	// }

	// t.Run("rectangles", func(t *testing.T) {
	// 	rectangle := ch5.Rectangle{5.0, 10.0}
	// 	checkArea(t, rectangle, 50.0)
	// })

	// t.Run("circles", func(t *testing.T) {
	// 	circle := ch5.Circle{10}
	// 	checkArea(t, circle, 314.1592653589793)
	// })
}
