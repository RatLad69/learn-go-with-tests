package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("bungholio")
	want := "Yoo, bungholio"

	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}
}
