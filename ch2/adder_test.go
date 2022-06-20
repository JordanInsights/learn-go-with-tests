package ch2_test

import (
	"learn-go-with-tests/ch2"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := ch2.Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}
