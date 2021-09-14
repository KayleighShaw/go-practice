package main

func Search(dictionary map[string]string, word string) string {
	// getting a value out of a map is the same as getting a value out of an array - map[key]
	return dictionary[word]
}
