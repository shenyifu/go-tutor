package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "world")
	got := buffer.String()
	want := "hello world"
	if got != want {
		t.Errorf("want %s got %s", want, got)

	}
}
