package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) {

		got := Hello("Raghul")

		want := "Hello, Raghul!"

		if got != want {
			t.Errorf("got %q and wanted %q", got, want)
		}

	})

	t.Run("Saying hello without parameters", func(t *testing.T) {

		got := Hello("")
		want := "Hello, Bangaari!"

		if got != want {
			t.Errorf("got %q and wanted %q", got, want)
		}
	})

}
