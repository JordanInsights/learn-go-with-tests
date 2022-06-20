package iteration_test

import (
	"learn-go-with-tests/iteration"
	"testing"
)

func TestIt(t *testing.T) {
	repeated := iteration.Repeat("a")
	expected := ""

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}
