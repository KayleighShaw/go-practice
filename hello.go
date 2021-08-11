package main

import "fmt"

// func main() {
// 	fmt.Println("Hello, world!")
// }

// Domain - good to seperate your "domain" code from the outside world (side-effects)
func Hello(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello("world"))
}
