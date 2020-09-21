package lcjune

import "testing"

func TestHIndex1(t *testing.T) {
	var h int
	var cit []int

	cit = []int{
		0, 1, 3, 5, 6,
	}

	h = hIndex1(cit)
	if h != 3 {
		t.Error(h, "should be", 3)
	}

	cit = []int{
		3, 4, 5, 8, 10,
	}

	h = hIndex1(cit)
	if h != 4 {
		t.Error(h, "should be", 4)
	}

	cit = []int{
		3, 3, 5, 8, 25,
	}

	h = hIndex1(cit)
	if h != 3 {
		t.Error(h, "should be", 3)
	}

	cit = []int{
		0, 0, 1,
	}

	h = hIndex1(cit)
	if h != 1 {
		t.Error(h, "should be", 1)
	}
}
