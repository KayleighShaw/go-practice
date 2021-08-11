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
	// the ":=" syntax lets us reuse some values in our test for readability
	got := Hello("Chris")
	want := "Hello, Chris"

	if got != want {
		// will return a string with the results for got and want, f stands for format
		t.Errorf("got %q want %q", got, want)
	}
}
