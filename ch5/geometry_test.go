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
		name    string
		shape   ch5.Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: ch5.Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: ch5.Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: ch5.Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})
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
