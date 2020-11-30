package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	if url == "123" {
		return true
	}
	return false
}

func mockSlowWebsiteChecker(url string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func TestCheckWebsite(t *testing.T) {
	website := []string{
		"123",
		"456",
		"789",
	}

	res := CheckWebsite(mockWebsiteChecker, website)
	exp := map[string]bool{
		"123": true,
		"456": false,
		"789": false,
	}

	if !reflect.DeepEqual(exp, res) {
		t.Fatalf("want %v got %v", exp, res)
	}
}

func BenchmarkWebsiteChecker(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "haha"
	}
	for i := 0; i < b.N; i++ {
		CheckWebsite(mockSlowWebsiteChecker, urls)
	}

}
