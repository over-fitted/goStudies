package testHelpers

import (
	"fmt"
	"testing"
)

func AssertCorrectMessage(t *testing.T, got, want interface{}) {
	t.Helper()
	if got != want {
		t.Errorf("expected %s, got %s", fmt.Sprintf("%v", want), fmt.Sprintf("%v", got))
	}
}
