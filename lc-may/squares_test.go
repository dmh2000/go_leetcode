package lcmay

import "testing"

func TestSquares1(t *testing.T) {
	var m [][]int
	var n int

	m = [][]int{
		{0, 1, 1, 1},
		{1, 1, 1, 1},
		{0, 1, 1, 1},
	}

	printSquares(m)
	n = countSquares(m)
	if n != 15 {
		t.Error(n, " should be ", 15)
	}
	n = countSquares(m)
	printSquares(m)
}
