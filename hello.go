package main

import "fmt"

const hello = "Hello, "

// Hello is
func Hello(name string) string {
	if name == "" {
		name = "world"
	}

	return hello + name
}

func main() {
	fmt.Println(Hello("name"))
}
