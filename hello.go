package main

import (
	"fmt"
	"net/http"
	g "./greeting"
)

const hello = "Hello, "
const hola = "Hola, "
const spanish = "spanish"
const bonjour = "Bonjour, "
const french = "french"

// Hello is
func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}

	return prefix(language) + name
}

func prefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = bonjour
	case spanish:
		prefix = hola
	default:
		prefix = hello
	}

	return
}

func main() {
	err := http.ListenAndServe(":5000", http.HandlerFunc(g.HandlerGreeting))

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(Hello("name", ""))
}
