package main

import "fmt"

// func main() {
// 	fmt.Println("Hello, world!")
// }

//constants improve performance + saves you creating the "Hello, " string every time the Hello func is called
const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefiex = "Hola, "
const frenchHelloPrefix = "Bonjour, "

// Domain - good to seperate your "domain" code from the outside world (side-effects)
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	if language == spanish {
		return spanishHelloPrefiex + name
	}
	if language == french {
		return frenchHelloPrefix + name
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
