package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) {

		got := Hello("Raghul", "English")
		want := "Hello, Raghul!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying hello without parameters", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, Bangaari!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying hello in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying hello in French", func(t *testing.T) {
		got := Hello("Fred", "French")
		want := "Bonjour, Fred!"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q and wanted %q", got, want)
	}
}
