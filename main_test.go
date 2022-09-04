package main

import (
	"testing"
)

func TestAddTest(test *testing.T) {
	if add(1, 2) != 3 {
		test.Errorf("add() = %v, want %v", add(1, 2), 3)
	}
}