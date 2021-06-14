package main

import "testing"

func TestHello(t *testing.T) {
	// accepts both T  for test and B for benchmark
	assertCorrectMessage := func(t testing.TB, got, want string) {
		// tells programme that this is a helper function
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello World' when empty string supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

}

func TestProverb(t *testing.T) {
	want := "Concurrency is not parallelism."
	if got := Proverb(); got != want {
		t.Errorf("Proverb() = %q, want %q", got, want)
	}
}
