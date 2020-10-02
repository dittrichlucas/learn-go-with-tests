package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func TestCheckWebsite(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	got := CheckWebsite(mockWebsiteChecker, websites)
	want := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"waat://furhurterwe.geds":    false,
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func mockWebsiteChecker(url string) bool {
	if url == "waat://furhurterwe.geds" {
		return false
	}

	return true
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarCheckWebsite(b *testing.B) {
	websites := make([]string, 100)

	for i := 0; i < len(websites); i++ {
		websites[i] = "an url"
	}

	for i := 0; i < b.N; i++ {
		CheckWebsite(slowStubWebsiteChecker, websites)
	}
}
