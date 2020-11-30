package select_my

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	slowUrl := makeDelayServer(20 * time.Millisecond)
	fastUrl := makeDelayServer(0 * time.Millisecond)

	defer fastUrl.Close()
	defer slowUrl.Close()

	want := fastUrl.URL
	got := Racer(slowUrl.URL, fastUrl.URL)
	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func makeDelayServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
