package select_my

import (
	"net/http"
)

func Racer(a, b string) (winner string) {
	select {
	case <-ping(a):
		return a
	case <-ping(b):
		return b
	}
}

func ping(url string) chan bool {
	channel := make(chan bool)
	go func() {
		http.Get(url)
		channel <- true
	}()
	return channel
}
