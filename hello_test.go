package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Ratman")
		want := "Yoo, Ratman"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Say 'Yoo, bungholio' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Yoo, bungholio"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
