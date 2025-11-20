package main

import "testing"

func TestDoMath(t *testing.T) {
	total := doMath(5, 5, add)

	if total != 10 {
		t.Errorf("Sum was incorrect, got %d, want %d", total, 10)
	}

	sub := doMath(42, 16, subtract)
	if sub != 26 {
		t.Errorf("Sum was incorrect, got %d, want %d", sub, 26)
	}
}
