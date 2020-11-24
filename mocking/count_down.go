package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

type ConfigableSleeper struct {
	duration time.Duration
}

func (c *ConfigableSleeper) Sleep() {
	time.Sleep(c.duration)
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func CountDown(buffer io.Writer, s Sleeper) {
	for i := 3; i > 0; i-- {
		s.Sleep()
		fmt.Fprintln(buffer, i)
	}
	s.Sleep()
	fmt.Fprintf(buffer, "go!")
}

func main() {
	sleeper := &ConfigableSleeper{1 * time.Second}
	CountDown(os.Stdout, sleeper)
}
