package runner

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRunner(t *testing.T) {
	t.Run("should return the server faster", func(t *testing.T) {
		slowServer := serverDelay(9 * time.Millisecond)
		fastServer := serverDelay(8 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		got, err := Runner(slowURL, fastURL)
		want := fastURL

		if err != nil {
			t.Fatalf("did not expect an error: %s", err)
		}

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	})

	t.Run("should return an error if the server does not respond within 10s", func(t *testing.T) {
		server := serverDelay(25 * time.Millisecond)

		defer server.Close()

		_, err := Setup(server.URL, server.URL, 20*time.Millisecond)

		if err == nil {
			t.Errorf("expected an error and did not return: %s", err)
		}
	})

}

func serverDelay(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
