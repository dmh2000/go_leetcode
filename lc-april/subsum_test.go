package lc

import "testing"

func TestSubSum(t *testing.T) {
	a := []int{1, 1, 1}
	x := subarraySum(a, 2)
	if x != 2 {
		t.Error(x, " should be ", 2)
	}

	a = []int{1, 3, 2, 3, 5}
	x = subarraySum(a, 5)
	if x != 4 {
		t.Error(x, " should be ", 3)
	}

	a = []int{1, 1, 1, 1, 1, 1, 1, 1, 1}
	x = subarraySum(a, 2)
	if x != 8 {
		t.Error(x, " should be ", 8)
	}

	a = []int{1, 1, 2, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	x = subarraySum(a, 2)
	if x != 31 {
		t.Error(x, " should be ", 31)
	}
}

func TestSubSum2(t *testing.T) {
	s := 16384
	a := []int{}
	for i := 0; i < s; i++ {
		a = append(a, 1)
	}
	x := subarraySum(a, 2)
	if x != s-1 {
		t.Error(x, " should be ", s-1)
	}

}
