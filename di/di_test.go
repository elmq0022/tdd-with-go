package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Aaron")

	got := buffer.String()
	want := "Hello, Aaron"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
