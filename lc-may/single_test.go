package lcmay

import "testing"

func TestSingle1(t *testing.T) {
	var a []int
	var r int

	a = []int{3, 3, 7, 7, 10, 11, 11}
	r = singleNonDuplicate(a)
	if r != 10 {
		t.Error(r, " should be ", 10)
	}

	a = []int{3, 3, 4, 7, 7, 11, 11}
	r = singleNonDuplicate(a)
	if r != 4 {
		t.Error(r, " should be ", 4)
	}

	a = []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6}
	r = singleNonDuplicate(a)
	if r != 6 {
		t.Error(r, " should be ", 6)
	}

	a = []int{0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5}
	r = singleNonDuplicate(a)
	if r != 0 {
		t.Error(r, " should be ", 0)
	}

	a = []int{0, 0, 1, 2, 2, 3, 3, 4, 4, 5, 5}
	r = singleNonDuplicate(a)
	if r != 1 {
		t.Error(r, " should be ", 1)
	}

	a = []int{0, 0, 1, 1, 2, 3, 3, 4, 4, 5, 5}
	r = singleNonDuplicate(a)
	if r != 2 {
		t.Error(r, " should be ", 2)
	}

	a = []int{0, 0, 1, 1, 2, 2, 3, 4, 4, 5, 5}
	r = singleNonDuplicate(a)
	if r != 3 {
		t.Error(r, " should be ", 3)
	}

	a = []int{0, 0, 1, 1, 2, 2, 3, 3, 4, 5, 5}
	r = singleNonDuplicate(a)
	if r != 4 {
		t.Error(r, " should be ", 4)
	}

	a = []int{0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5}
	r = singleNonDuplicate(a)
	if r != 5 {
		t.Error(r, " should be ", 5)
	}

}
