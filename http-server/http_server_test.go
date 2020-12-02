package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

/**
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}**/

func TestPlayerServer(t *testing.T) {
	server := &PlayServer{&StubPlayerStore{map[string]int{
		"haha": 20,
		"baba": 10,
	}}}
	t.Run("test haha", func(t *testing.T) {
		request := newGetScoreRequest("haha")

		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertEqual(t, response.Body.String(), "20")
	})

	t.Run("test baba", func(t *testing.T) {
		request := newGetScoreRequest("baba")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusOK)
		assertEqual(t, response.Body.String(), "10")

	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		request := newGetScoreRequest("Apollo")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Code
		want := http.StatusNotFound

		if got != want {
			t.Errorf("got status %d want %d", got, want)
		}
	})

}

func TestStoreWins(t *testing.T) {

}

func assertEqual(t *testing.T, got, want string) {
	if got != want {
		t.Errorf("nomber error want '%s' got %s", want, got)
	}
}

func assertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}

func newGetScoreRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, "/players/"+name, nil)
	return request
}
