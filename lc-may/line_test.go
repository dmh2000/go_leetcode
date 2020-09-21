package lcmay

import (
	"fmt"
	"testing"
)

func TestLine1(t *testing.T) {
	var b bool
	var a [][]int

	a = [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 5},
		{5, 6},
		{6, 7},
	}
	b = checkStraightLine(a)
	if !b {
		t.Error(b, " should be ", true)
	}
	fmt.Println("----------------------------------------")
	a = [][]int{
		{1, 2},
		{2, 4},
		{3, 4},
		{4, 5},
		{5, 6},
		{6, 7},
	}
	b = checkStraightLine(a)
	if b {
		t.Error(b, " should be ", false)
	}

	fmt.Println("----------------------------------------")
	a = [][]int{
		{1, 1},
		{2, 2},
		{3, 4},
		{4, 5},
		{5, 6},
		{6, 7},
	}
	b = checkStraightLine(a)
	if b {
		t.Error(b, " should be ", false)
	}

	fmt.Println("----------------------------------------")
	a = [][]int{
		{1, 1},
		{2, 1},
		{3, 1},
		{4, 1},
		{5, 1},
		{6, 1},
	}
	b = checkStraightLine(a)
	if !b {
		t.Error(b, " should be ", true)
	}

	fmt.Println("----------------------------------------")
	a = [][]int{
		{1, 1},
		{1, 2},
		{1, 3},
		{1, 4},
		{1, 5},
		{1, 6},
	}
	b = checkStraightLine(a)
	if !b {
		t.Error(b, " should be ", true)
	}

}
