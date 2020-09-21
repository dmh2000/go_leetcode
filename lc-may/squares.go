package lcmay

import "fmt"

type sq struct {
	row int
	col int
}

func printSquares(m [][]int) {
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			fmt.Printf("%2v ", m[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

func squaresRecursive(m [][]int, row int, col int, sz int, count *int) {
	var rlen int
	var clen int

	// number of columns to check
	clen = sz
	if col+clen > len(m[0]) {
		return
	}

	// number of rows to check
	rlen = sz
	if sz < rlen {
		rlen = sz
	}
	if row+rlen > len(m) {
		return
	}

	for r := row; r < row+rlen; r++ {
		for c := col; c < col+clen; c++ {
			if m[r][c] != 1 {
				// not a square
				return
			}
		}
	}

	*count++

	// TAIL RECURSION, MAKE IT ITERATIVE!
	squaresRecursive(m, row, col, sz+1, count)
}

func squaresIterative(m [][]int, row int, col int) int {
	var rlen int
	var clen int
	var lim int
	var sz int
	var count int

	// get max square size for this matrix
	lim = len(m[0])
	if lim < len(m) {
		lim = len(m)
	}

	count = 0
	sz = 1
outer:
	for sz < lim {
		// number of columns to check
		clen = sz
		// too big?
		if col+clen > len(m[0]) {
			break
		}

		// number of rows to check
		rlen = sz
		if sz < rlen {
			rlen = sz
		}
		// too big?
		if row+rlen > len(m) {
			break
		}

		for r := row; r < row+rlen; r++ {
			for c := col; c < col+clen; c++ {
				if m[r][c] != 1 {
					// not a square
					break outer
				}
			}
		}

		count++

		// next size
		sz++
	}
	return count
}

func countSquares(m [][]int) int {
	var count int

	count = 0
	for r := 0; r < len(m); r++ {
		for c := 0; c < len(m[0]); c++ {
			if m[r][c] != 0 {
				count += squaresIterative(m, r, c)
				// squaresRecursive(m, r, c, 1, &count)
			}
		}
	}

	return count
}
