package concurrency

// WebsiteChecker is
type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

// CheckWebsite is
func CheckWebsite(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	channel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			channel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		result := <-channel
		results[result.string] = result.bool
	}

	return results
}
