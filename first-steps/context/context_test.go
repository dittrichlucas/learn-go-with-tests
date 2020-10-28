package context

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHandler(t *testing.T) {
	data := "Hello, world!"

	t.Run("test 1", func(t *testing.T) {
		store := &SpyStore{response: data, t: t}
		svr := Server(store)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		svr.ServeHTTP(res, req)

		if res.Body.String() != data {
			t.Errorf("got '%s', want '%s'", res.Body.String(), data)
		}
	})

	t.Run("should cancel work after 100 milliseconds", func(t *testing.T) {
		store := &SpyStore{response: data, t: t}
		svr := Server(store)

		req := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(req.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		req = req.WithContext(cancellingCtx)

		res := &SpyResponseWriter{}

		svr.ServeHTTP(res, req)

		if res.written {
			t.Errorf("an answer should not be written")
		}
	})

}

type SpyResponseWriter struct {
	written bool
}

type SpyStore struct {
	response string
	t        *testing.T
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, key := range s.response {
			select {
			case <-ctx.Done():
				s.t.Log("spy store was cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(key)
			}
		}
		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implement")
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}

// func (s *SpyStore) Cancel() {
// 	s.cancelled = true
// }

// func (s *SpyStore) assertWasCancelled() {
// 	s.t.Helper()
// 	if !s.cancelled {
// 		s.t.Errorf("ops")
// 	}
// }

// func (s *SpyStore) assertWasNotCancelled() {
// 	s.t.Helper()
// 	if s.cancelled {
// 		s.t.Errorf("")
// 	}
// }
