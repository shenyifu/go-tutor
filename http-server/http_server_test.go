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
	},
		nil,
	}}
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
	store := StubPlayerStore{
		map[string]int{}, nil,
	}
	server := &PlayServer{&store}

	t.Run("it return accept on post", func(t *testing.T) {
		player := "Pepper"

		request := newPostWinRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)

		if len(store.winCalls) != 1 {
			t.Fatalf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
		}

		if store.winCalls[0] != player {
			t.Errorf("did not store correct winner got '%s' want '%s'", store.winCalls[0], player)
		}
	})
}

func TestRecordingWinsRetrivingThem(t *testing.T) {
	store := NewInMemoryPlayerStore()
	server := PlayServer{store}
	player := "haha"
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetScoreRequest(player))
	assertStatus(t, response.Code, http.StatusOK)

	assertEqual(t, response.Body.String(), "3")
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

func newPostWinRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodPost, "/players/"+name, nil)
	return request
}
