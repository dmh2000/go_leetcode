package lc

import "testing"

func TestProduct1(t *testing.T) {
	a := []int{2, 3, 4, 5}
	b := productExceptSelf(a)
	c := []int{60, 40, 30, 24}
	for i := 0; i < len(b); i++ {
		if b[i] != c[i] {
			t.Errorf("%v:%v should be equal to %v", i, b[i], c[i])
		}
	}
}

func TestProduct2(t *testing.T) {
	a := []int{1, 2}
	b := productExceptSelf(a)
	c := []int{2, 1}
	for i := 0; i < len(b); i++ {
		if b[i] != c[i] {
			t.Errorf("%v:%v should be equal to %v", i, b[i], c[i])
		}
	}
}
