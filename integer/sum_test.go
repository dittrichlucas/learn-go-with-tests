package integer

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("should return a sum of slice", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("\ngot: %d\nwant: %d\ngiven: %v", got, want, numbers)
		}
	})
}

func ExampleSum() {
	numbers := []int{1, 2, 3, 4, 5}
	sum := Sum(numbers)
	fmt.Println(sum)
	// Output: 15
}
