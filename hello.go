package main

import "fmt"

// func main() {
// 	fmt.Println("Hello, world!")
// }

//constants improve performance + saves you creating the "Hello, " string every time the Hello func is called
const spanish = "Spanish"
const french = "French"
const japanese = "Japanese"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefiex = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const japaneseHelloPrefix = "こんにちは, "

// Domain - good to seperate your "domain" code from the outside world (side-effects)
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

//named return value, creates a prefix variable assigned (0/"" - type dependant)
//function is lowercase = private, won't show the inner workings to the world
func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefiex
	case japanese:
		prefix = japaneseHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return //returns the named variable just with "return"
}

func main() {
	fmt.Println(Hello("world", ""))
}
