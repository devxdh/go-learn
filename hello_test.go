package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Dev")
	want := "Hello, Dev"

	if got != want {
		t.Errorf("got %q wnat %q", got, want)
	}
}
