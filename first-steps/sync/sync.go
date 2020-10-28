package sync

import "sync"

// Count is
type Count struct {
	mu    sync.Mutex
	value int
}

// NewCount is
func NewCount() *Count {
	return &Count{}
}

// Increment is
func (c *Count) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Value is
func (c *Count) Value() int {
	return c.value
}
