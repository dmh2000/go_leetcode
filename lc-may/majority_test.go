package lcmay

import "testing"

func TestMajority1(t *testing.T) {
	a := []int{
		2, 2, 1, 1, 1, 2, 2,
	}

	n := majorityElement(a)
	if n != 2 {
		t.Error(n, " should be ", 2)
	}

	a = []int{
		3, 2, 3,
	}

	n = majorityElement(a)
	if n != 3 {
		t.Error(n, " should be ", 3)
	}
}
