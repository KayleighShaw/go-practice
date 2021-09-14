package main

import (
	"testing"
)

func TestSearch(t *testing.T) {
	// declare a map with the map keyword, requires the [key] type and the value type after []
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := "could not find the word you were looking for"

		// also protecting assertStrings to ensure we don't called .Error() on nil
		if err == nil {
			t.Fatal("expected to get an error.")
		}

		// errors can be converted to a string with the .Error() method
		assertStrings(t, err.Error(), want)
	})

}

// // helps make the implementation more general
func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
