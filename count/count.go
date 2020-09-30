package count

import (
	"fmt"
	"io"
	"time"
)

// Sleeper is
type Sleeper interface {
	Pause()
}

// SleeperDefault is
type SleeperDefault struct {
	duration time.Duration
	pause    func(time.Duration)
}

// Pause is
func (t *SleeperDefault) Pause() {
	t.pause(t.duration)
}

const lastWord = "Go!"
const start = 3

// Count is
func Count(output io.Writer, sleeper Sleeper) {
	for i := start; i > 0; i-- {
		sleeper.Pause()
		fmt.Fprintln(output, i)
	}

	sleeper.Pause()
	fmt.Fprint(output, lastWord)
}
