package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) {

		got := Hello("Raghul")
		want := "Hello, Raghul!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying hello without parameters", func(t *testing.T) {

		got := Hello("")
		want := "Hello, Bangaari!"
		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q and wanted %q", got, want)
	}
}
