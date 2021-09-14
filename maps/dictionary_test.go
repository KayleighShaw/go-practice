package main

import "testing"

func TestSearch(t *testing.T) {
	// declare a map with the map keyword, requires the [key] type and the value type after []
	dictionary := Dictionary{"test": "this is just a test"}

	got := dictionary.Search("test")
	want := "this is just a test"

	assertString(t, got, want)
}

// helps make the implementation more general
func assertString(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
