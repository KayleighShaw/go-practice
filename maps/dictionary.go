package main

// Dictoinary type acts as a thin wrapper around map
type Dictionary map[string]string

// with the custom type defined we can create the Search method
func (d Dictionary) Search(word string) string {
	return d[word]
}
