package main

import "errors"

// Dictoinary type acts as a thin wrapper around map
type Dictionary map[string]string

// with the custom type defined we can create the Search method
func (d Dictionary) Search(word string) (string, error) {
	// map can return 2 values. The second value is a boolean which indicates if a key was found successfully
	// this allows us to differentiate between a word that doesn't exist vs a word that doesn't have a definition
	definition, ok := d[word]
	if !ok {
		return "", errors.New("could not find the word you were looking for")
	}
	return definition, nil
}
