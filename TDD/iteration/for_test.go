package iteration

import (
	"TDD/testHelpers"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("default iterations", func(t *testing.T) {
		got, err := Repeat("a")
		want := "aaaaa"
		if err != nil {
			t.Error(err)
		}
		testHelpers.AssertCorrectMessage(t, got, want)
	})
	t.Run("2 iterations", func(t *testing.T) {
		got, err := Repeat("a", 2)
		want := "aa"
		if err != nil {
			t.Error(err)
		}
		testHelpers.AssertCorrectMessage(t, got, want)
	})
}

func BenchmarkRepeat(b *testing.B) {
	// b.N benchmarks high load
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
