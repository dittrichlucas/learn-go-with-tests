package main

import "testing"

func TestHello(t *testing.T) {
	error := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("\ngot: '%s'\nwant: '%s'", got, want)
		}
	}

	t.Run("should say hello", func(t *testing.T) {
		got := Hello("Lucas")
		want := "Hello, Lucas"
		error(t, got, want)
	})

	t.Run("should say 'Hello, world' when a empty value is passed", func(t *testing.T) {
		got := Hello("")
		want := "Hello, world"
		error(t, got, want)
	})

}
