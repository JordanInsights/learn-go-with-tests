package ch2_test

import (
	"fmt"
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

func ExampleAdd() {
	sum := ch2.Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
