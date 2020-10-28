package runner

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondsLimit = 10 * time.Second

// Runner is
func Runner(first, second string) (winner string, err error) {
	return Setup(first, second, tenSecondsLimit)
}

// Setup is
func Setup(first, second string, limit time.Duration) (winner string, err error) {
	select {
	case <-ping(first):
		return first, nil
	case <-ping(second):
		return second, nil
	case <-time.After(limit):
		return "", fmt.Errorf("tempo excedido para %s e %s", first, second)
	}
}

func ping(URL string) chan bool {
	channel := make(chan bool)

	go func() {
		http.Get(URL)
		channel <- true
	}()

	return channel
}
