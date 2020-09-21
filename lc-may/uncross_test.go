package lcmay

import (
	"fmt"
	"testing"
)

func TestUncross(t *testing.T) {
	var a []int
	var b []int
	var r int

	a = []int{1, 2, 3, 4}
	b = []int{1, 4, 5, 6}
	r = maxUncrossedLines(a, b)
	if r != 2 {
		t.Error(r, "should be", 2)
	}
	fmt.Printf("%v \n\n", r)

	a = []int{1, 4, 5, 6}
	b = []int{1, 2, 3, 4}
	r = maxUncrossedLines(a, b)
	if r != 2 {
		t.Error(r, "should be", 2)
	}
	fmt.Printf("%v \n\n", r)

	a = []int{2, 5, 1, 2, 5}
	b = []int{10, 5, 2, 1, 5, 2}
	r = maxUncrossedLines(a, b)
	if r != 3 {
		t.Error(r, "should be", 3)
	}
	fmt.Printf("%v \n\n", r)

	a = []int{1, 3, 7, 1, 7, 5}
	b = []int{1, 9, 2, 5, 1}
	r = maxUncrossedLines(a, b)
	if r != 2 {
		t.Error(r, "should be", 2)
	}
	fmt.Printf("%v \n\n", r)

	a = []int{1, 1, 2, 1, 2}
	b = []int{1, 3, 2, 3, 1}
	r = maxUncrossedLines(a, b)
	if r != 3 {
		t.Error(r, "should be", 3)
	}
	fmt.Printf("%v \n\n", r)

}
