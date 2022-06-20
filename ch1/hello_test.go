package ch1_test

import (
	"learn-go-with-tests/ch1"
	"testing"
)

func TestHello(t *testing.T) {
	got := ch1.Hello("Jordan")
	want := "Hello, Jordan"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
