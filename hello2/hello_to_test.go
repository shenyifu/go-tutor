package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMsg := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%q' want '%q' ", got, want)
		}
	}
	t.Run("say hello", func(t *testing.T) {
		got := Hello("Charis")
		want := "Hello,Charis"
		assertCorrectMsg(t, got, want)
	})

	t.Run("say default", func(t *testing.T) {
		got := Hello("")
		want := "Hello,world"
		assertCorrectMsg(t, got, want)
	})
}
