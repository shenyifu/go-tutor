package concurrency

type WebsiteChecker func(string) bool

type result struct {
	url   string
	valid bool
}

func CheckWebsite(wc WebsiteChecker, urls []string) map[string](bool) {
	presult := make(map[string]bool)
	channel := make(chan result)
	for _, url := range urls {
		go func(url string) {
			channel <- result{url, wc(url)}
		}(url)
	}
	for i := 0; i < len(urls); i++ {
		result := <-channel
		presult[result.url] = result.valid
	}
	return presult
}
