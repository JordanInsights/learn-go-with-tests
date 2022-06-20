package iteration_test

import (
	"learn-go-with-tests/iteration"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := iteration.Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		iteration.Repeat("a")
	}
}
