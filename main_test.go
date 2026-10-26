package main

import "testing"

func TestHello(t *testing.T) {
	want := "Hello Go"
	if got := hello(); got != want {
		t.Fatalf("want %q, got %q", want, got)
	}
}
