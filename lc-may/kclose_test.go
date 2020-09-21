package lcmay

import (
	"fmt"
	"testing"
)

func TestKClose1(t *testing.T) {
	var a [][]int
	var p [][]int

	p = [][]int{
		{1, 3},
		{-2, 2},
	}
	a = kClosest(p, 1)
	fmt.Println(a)

	p = [][]int{
		{3, 3},
		{5, -1},
		{-2, 4},
	}
	a = kClosest(p, 2)
	fmt.Println(a)
}
