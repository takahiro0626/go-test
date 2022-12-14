package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	t.Run("plus number", func(t *testing.T) {
		got := add(1, 2)
		want := 3
		if got != want {
			t.Errorf("add() = %v, want %v", add(1, 2), 3)
		}

	})

	t.Run("minus number", func(t *testing.T) {
		got := add(-1, -2)
		want := -3
		if got != want {
			t.Errorf("add() = %v, want %v", add(-1, -2), -3)
		}
	})

	t.Run("both", func(t *testing.T) {
		got := add(1, -2)
		want := -1
		if got != want {
			t.Errorf("add() = %v, want %v", add(1, -2), -1)
		}
	})

}
func TestAddAssert(t *testing.T) {
	assert.Equal(t, add(1, 1), 2)
	assert.NotEqual(t, add(1, 1), 1)
}