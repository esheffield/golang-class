package main

import "testing"

func TestAdd(t *testing.T) {
	v := add(12, 16)

	if v != 28 {
		t.Error("Expected 28, got ", v)
	}
}
