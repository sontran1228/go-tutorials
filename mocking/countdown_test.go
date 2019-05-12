package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}

	Countdown(buffer, spySleeper)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}

	if spySleeper.Calls != 4 {
		t.Errorf("not enough calls to sleeper, want 4 got %d", spySleeper.Calls)
	}
}

// SpySleeper is the spy to mock the "time.Sleep"
// Spies are a kind of mock which can record how a dependency is used.
// They can record the arguments sent in, how many times, etc.
// In our case, we're keeping track of how many times Sleep() is called so we can check it in our test.
type SpySleeper struct {
	Calls int
}

// the mocking method "Sleep", don't need to wait for 4s to finish one testcase.
func (s *SpySleeper) Sleep() {
	s.Calls++
}
