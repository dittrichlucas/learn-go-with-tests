package main

import "fmt"

// Hello is
func Hello(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello("name"))
}
