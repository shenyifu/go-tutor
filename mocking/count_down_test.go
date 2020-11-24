package main

import (
	"bytes"
	"reflect"
	"testing"
)

type CountDownOperation struct {
	Calls []string
}

func (c *CountDownOperation) Sleep() {
	c.Calls = append(c.Calls, "sleep")
}

func (c *CountDownOperation) Write(p []byte) (n int, err error) {
	c.Calls = append(c.Calls, "write")
	return
}
func TestCountDown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}

	CountDown(buffer, spySleeper)
	want := `3
2
1
go!`

	if spySleeper.Calls != 4 {
		t.Errorf("call time error")
	}

	if want != buffer.String() {
		t.Errorf("got %s want %s", buffer.String(), want)
	}

	t.Run("test order", func(t *testing.T) {
		spySleepPrinter := &CountDownOperation{}
		CountDown(spySleepPrinter, spySleepPrinter)
		want := []string{
			"sleep",
			"write",
			"sleep",
			"write",
			"sleep",
			"write",
			"sleep",
			"write",
		}
		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("want %v goy %v", want, spySleepPrinter.Calls)
		}
	})

}
