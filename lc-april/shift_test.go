package lc

import "testing"

func TestShift(t *testing.T) {
	var s string

	a := [][]int{
		{0, 1}, {1, 2},
	}
	s = stringShift("abc", a)
	if s != "cab" {
		t.Error(s, " should be ", "cab")
	}

	b := [][]int{
		{1, 1}, {1, 1}, {0, 2}, {1, 3},
	}
	s = stringShift("abcdefg", b)
	if s != "efgabcd" {
		t.Error(s, " should be ", "efgabcd")
	}
}
