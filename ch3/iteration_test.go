package ch3_test

import (
	"fmt"
	"learn-go-with-tests/ch3"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := ch3.Repeat("a", 3)
	expected := "aaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ch3.Repeat("a", 3)
	}
}

func ExampleRepeat() {
	s := ch3.Repeat("a", 5)
	fmt.Println(s)
	// Output: aaaaa
}
