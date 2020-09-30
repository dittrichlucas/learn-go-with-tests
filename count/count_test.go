package count

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCount(t *testing.T) {
	t.Run("print 3 until Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}

		Count(buffer, &CounterSpy{})

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	})

	t.Run("pause before each print", func(t *testing.T) {
		sleeperSpy := &CounterSpy{}
		Count(sleeperSpy, sleeperSpy)

		want := []string{
			pause,
			write,
			pause,
			write,
			pause,
			write,
			pause,
			write,
		}

		if !reflect.DeepEqual(want, sleeperSpy.Calls) {
			t.Errorf("got '%v', want '%s'", sleeperSpy.Calls, want)
		}
	})

}

func TestSleeperDefault(t *testing.T) {
	pauseTime := 5 * time.Second

	timeSpy := &TimeSpy{}
	sleeper := SleeperDefault{pauseTime, timeSpy.Pausa}
	sleeper.Pause()

	if timeSpy.pauseDuration != pauseTime {
		t.Errorf("got '%v', want '%v'", timeSpy.pauseDuration, pauseTime)
	}
}

// CounterSpy is
type CounterSpy struct {
	Calls []string
}

// TimeSpy is
type TimeSpy struct {
	pauseDuration time.Duration
}

// Pause is
func (c *CounterSpy) Pause() {
	c.Calls = append(c.Calls, pause)
}

// Write is
func (c *CounterSpy) Write(p []byte) (n int, err error) {
	c.Calls = append(c.Calls, write)
	return
}

func (t *TimeSpy) Pausa(duration time.Duration) {
	t.pauseDuration = duration
}

const pause = "pause"
const write = "write"
