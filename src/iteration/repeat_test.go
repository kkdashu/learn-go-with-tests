package iteration

import (
  "testing"

  "github.com/stretchr/testify/assert"
)

func TestRepeat(t *testing.T) {
  repeated := Repeat("a")
  expected := "aaaaa"

  assert.Equal(t, repeated, expected)
}

func BenchmarkRepeat(b *testing.B) {
  for i := 0; i < b.N; i++ {
    Repeat("a")
  }
}
