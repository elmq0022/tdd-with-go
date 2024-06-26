package main

import (
	"testing"
)

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say  'Hello, World'when an empty string is provided", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Marcie", "French")
		want := "Bonjour, Marcie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Norwegian", func(t *testing.T) {
		got := Hello("Bjorn", "Norwegian")
		want := "Hei, Bjorn"
		assertCorrectMessage(t, got, want)
	})
}
