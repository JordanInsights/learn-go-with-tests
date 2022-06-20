package iteration_test

import (
	"fmt"
	"learn-go-with-tests/iteration"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := iteration.Repeat("a", 3)
	expected := "aaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		iteration.Repeat("a", 3)
	}
}

func ExampleRepeat() {
	s := iteration.Repeat("a", 5)
	fmt.Println(s)
	// Output: aaaaa
}
