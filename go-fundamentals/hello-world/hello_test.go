package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, world"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHelloFriend(t *testing.T) {
	got := HelloFriend("Joe")
	want := "Hello, Joe"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
