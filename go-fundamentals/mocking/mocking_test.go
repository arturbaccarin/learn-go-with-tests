package mocking

import (
	"bytes"
	"testing"
)

// If we can mock time.Sleep we can use dependency injection to use it instead of a "real"
// time.Sleep and then we can spy on the calls to make assertions on them.

// It's an important skill to be able to slice up requirements as small
// as you can so you can have working software.
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
		t.Errorf("got %q want %q", got, want)
	}

	if spySleeper.Calls != 3 {
		t.Errorf("not enough calls to sleeper, want 3 got %d", spySleeper.Calls)
	}
}
