package main

import (
	"log"
	"net/http"
)

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

type InMemoryPlayerStore struct {
	store map[string]int
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

func main() {
	handler := &PlayServer{NewInMemoryPlayerStore()}
	if err := http.ListenAndServe(":5000", handler); err != nil {
		log.Fatalf("could not listen to port %v", err)
	}
}
