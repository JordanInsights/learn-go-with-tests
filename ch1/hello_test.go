package ch1_test

import (
	"learn-go-with-tests/ch1"
	"testing"
)

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := ch1.Hello("Jordan")
		want := "Hello, Jordan"

		assertCorrectMessage(t, got, want)
	})

	t.Run("emptry string defaults to 'World'", func(t *testing.T) {
		got := ch1.Hello("")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
}
