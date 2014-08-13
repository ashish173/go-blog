package main

import "testing"

func TestAdd(t *testing.T) {
	var v int
	v = Add([]int{1, 2})
	if v != 3 {
		t.Error("Expected 1.5, got ", v)
	}
}
