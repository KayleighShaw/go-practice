package main

import "fmt"

// func main() {
// 	fmt.Println("Hello, world!")
// }

//constants improve performance + saves you creating the "Hello, " string every time the Hello func is called
const englishHelloPrefix = "Hello, "

// Domain - good to seperate your "domain" code from the outside world (side-effects)
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
