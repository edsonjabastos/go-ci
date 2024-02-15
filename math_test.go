package main

import "testing"

func TestAdd(t *testing.T) {
	total := add(20, 20)
	if total != 40 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 40)
	}
}
