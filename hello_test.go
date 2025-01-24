package main

import "testing"

func TestHello(t *testing.T) {

        got := Hello("Raghul")

        want := "Hello, Raghul!"

        if got != want {
                t.Errorf("got %q and wanted %q", got, want)
        }

}
