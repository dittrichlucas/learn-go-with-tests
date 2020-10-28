package greeting

import (
	"fmt"
	"io"
	"net/http"
)

// Greeting is
func Greeting(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

// HandlerGreeting is
func HandlerGreeting(w http.ResponseWriter, r *http.Request) {
	Greeting(w, "world")
}
