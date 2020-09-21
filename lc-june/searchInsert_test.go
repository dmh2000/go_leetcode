package lcjune

import (
	"testing"
)

func TestSearchInsert1(t *testing.T) {
	var a []int
	var index int

	a = []int{1, 3, 5, 6}

	index = searchInsert(a, 5)
	if index != 2 {
		t.Error(index, " should be ", 2)
	}

	index = searchInsert(a, 2)
	if index != 1 {
		t.Error(index, " should be ", 1)
	}

	index = searchInsert(a, 0)
	if index != 0 {
		t.Error(index, " should be ", 0)
	}

	index = searchInsert(a, 7)
	if index != 4 {
		t.Error(index, " should be ", 4)
	}
}

func TestSearchInsert2(t *testing.T) {
	var a []int
	var index int

	a = []int{1}

	index = searchInsert(a, 0)
	if index != 0 {
		t.Error(index, " should be ", 0)
	}

	index = searchInsert(a, 2)
	if index != 1 {
		t.Error(index, " should be ", 1)
	}
}

func TestSearchInsert3(t *testing.T) {
	var a []int
	var index int

	a = []int{1, 3}

	index = searchInsert(a, 0)
	if index != 0 {
		t.Error(index, " should be ", 0)
	}

	index = searchInsert(a, 1)
	if index != 0 {
		t.Error(index, " should be ", 0)
	}

	index = searchInsert(a, 2)
	if index != 1 {
		t.Error(index, " should be ", 1)
	}

	index = searchInsert(a, 3)
	if index != 1 {
		t.Error(index, " should be ", 1)
	}

	index = searchInsert(a, 4)
	if index != 2 {
		t.Error(index, " should be ", 2)
	}
}

func TestSearchInsert4(t *testing.T) {
	var a []int
	var index int

	a = []int{1, 3, 5}

	index = searchInsert(a, 4)
	if index != 2 {
		t.Error(index, " should be ", 2)
	}

}
