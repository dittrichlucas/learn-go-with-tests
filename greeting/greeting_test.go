package greeting

import (
	"bytes"
	"testing"
)

func TestGreeting(t *testing.T) {
	buffer := bytes.Buffer{}
	Greeting(&buffer, "Lucas")

	got := buffer.String()
	want := "Hello, Lucas"

	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}
