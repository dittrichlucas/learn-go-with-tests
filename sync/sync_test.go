package sync

import (
	"sync"
	"testing"
)

func TestCountSync(t *testing.T) {
	t.Run("should increment a count 3 times", func(t *testing.T) {
		count := NewCount()
		count.Increment()
		count.Increment()
		count.Increment()

		checkCount(t, count, 3)
	})

	t.Run("should run concurrently and safe", func(t *testing.T) {
		wantCount := 1000
		count := NewCount()

		var wg sync.WaitGroup
		wg.Add(wantCount)

		for i := 0; i < wantCount; i++ {
			go func(w *sync.WaitGroup) {
				count.Increment()
				w.Done()
			}(&wg)
		}
		wg.Wait()

		checkCount(t, count, wantCount)
	})
}

func checkCount(t *testing.T, result *Count, want int) {
	t.Helper()
	if result.Value() != want {
		t.Errorf("got %d, want %d", result.Value(), want)
	}
}
