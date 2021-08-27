package main

import (
	"testing"
)

func TestFun(t *testing.T) {
	out := validTree(5, [][]int{{0, 1}, {1, 2}, {2, 3}, {1, 3}, {1, 4}})
	want := false
	if out != want {
		t.Errorf("got %t, want %t", out, want)
	}
}
