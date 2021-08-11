package main

/*
Rules for Go Testing:
- Needs to be a file like xxxx_test.go
- Test function must start with word Test
- The test function takes one argument only t *testing.T
	- your t of type *testing.T is a "hook" into the testing framework
- Need to import "testing"
- Lauch docs locally: godoc -http :8000
*/

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		// the ":=" syntax lets us reuse some values in our test for readability
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Faustine", "French")
		want := "Bonjour, Faustine"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Japanese", func(t *testing.T) {
		got := Hello("Mai", "Japanese")
		want := "こんにちは, Mai"
		assertCorrectMessage(t, got, want)
	})
}
