package main

import "testing"

func TestCount(t *testing.T) {
	if got := Count(100); got == 0 {
		t.Fatalf("Count(100) = 0")
	}
}
