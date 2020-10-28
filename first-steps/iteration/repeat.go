package iteration

import "strings"

// Repeat receive a string and return a string
func Repeat(char string, amount int) string {
	return strings.Repeat(char, amount)
}
