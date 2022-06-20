package ch5_test

import (
	"learn-go-with-tests/ch5"
	"testing"
)

func assertEquality(t *testing.T, got, want float64) {
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestPerimeter(t *testing.T) {
	got := ch5.Perimeter(5.0, 10.0)
	want := 30.0

	assertEquality(t, got, want)
}

func TestArea(t *testing.T) {
	got := ch5.Area(5.0, 10.0)
	want := 50.0

	assertEquality(t, got, want)
}
