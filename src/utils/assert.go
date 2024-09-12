package utils

import (
  "testing"
)

func AssertEqual(t testing.TB, got, want string) {
  t.Helper()
  if got != want {
    t.Errorf("got %q want %q", got, want)
  }
}
