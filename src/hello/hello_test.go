package main

import (
	"testing"

  "github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	got := Hello("World", "")
	want := "Hello, World!"
	assert.Equal(t, got, want)

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris!"
		assert.Equal(t, want, got)
	})

	t.Run("say 'Hello, World' default", func(t *testing.T) {
		assert.Equal(t, "Hello, World!", Hello("", ""))
	})

  t.Run("in Spanish", func(t *testing.T) {
    got := Hello("Elodie", "Spanish")
    want := "Hola, Elodie!"
    assert.Equal(t, got, want)
  })
}
