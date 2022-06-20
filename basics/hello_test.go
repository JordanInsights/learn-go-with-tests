package basics_test

import (
	"learn-go-with-tests/basics"
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
		got := basics.Hello("Jordan", "")
		want := "Hello, Jordan"

		assertCorrectMessage(t, got, want)
	})

	t.Run("emptry string defaults to 'World'", func(t *testing.T) {
		got := basics.Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := basics.Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := basics.Hello("Napolean", "French")
		want := "Bonjour, Napolean"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in German", func(t *testing.T) {
		got := basics.Hello("Herman", "German")
		want := "Guten Tag, Herman"
		assertCorrectMessage(t, got, want)
	})
}
