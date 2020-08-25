package integer

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	got := Sum(2, 2)
	want := 4

	if got != want {
		t.Errorf("\ngot: %d\nwant: %d", got, want)
	}
}

func ExampleSum() {
	sum := Sum(1, 5)
	fmt.Println(sum)
	// Output: 6
}
