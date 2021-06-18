package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	// bytes.Buffer implements io.Writer, is a testable version
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
