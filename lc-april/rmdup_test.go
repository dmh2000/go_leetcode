package lc

import (
	"fmt"
	"testing"
)

func TestRmDup1(t *testing.T) {
	a := []int{1, 1, 2}
	i := removeDuplicates(a)
	if i != 2 {
		t.Error(i, " should be ", 2)
	}

	a = []int{1, 2, 2}
	i = removeDuplicates(a)
	if i != 2 {
		t.Error(i, " should be ", 2)
	}

	a = []int{0, 0, 1, 1, 2, 2, 3, 3, 4, 4}
	i = removeDuplicates(a)
	if i != 5 {
		t.Error(i, " should be ", 5)
	}
	fmt.Println(a)
}
