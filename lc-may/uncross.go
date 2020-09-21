package lcmay

import "fmt"

// if match ,previous diagonal
// else max of previous column and previous row
//

func printM(m [][]int) {
	for r := 0; r < len(m); r++ {
		for c := 0; c < len(m[0]); c++ {
			fmt.Printf("%2v ", m[r][c])
		}
		fmt.Println()
	}
}

func maxUncrossedLines(A []int, B []int) int {
	cols := len(A) + 1
	rows := len(B) + 1

	var m [][]int
	// var n [][]int

	m = make([][]int, rows)
	// n = make([][]int, rows)
	for i := 0; i < rows; i++ {
		m[i] = make([]int, cols)
		// n[i] = make([]int, cols)
	}

	// // populate n matrix with values
	// for i := 1; i < rows; i++ {
	// 	n[i][0] = B[i-1]
	// }
	// for i := 1; i < cols; i++ {
	// 	n[0][i] = A[i-1]
	// }

	// at all row/col
	// if match put previous diagonal + 1
	// else max of previous column and previous row
	for r := 1; r < rows; r++ {
		a := B[r-1]
		for c := 1; c < cols; c++ {
			b := A[c-1]
			if a == b {
				m[r][c] = 1 + m[r-1][c-1]
				// n[r][c] = a
			} else {
				x := m[r-1][c]
				y := m[r][c-1]
				if x > y {
					m[r][c] = x
				} else {
					m[r][c] = y
				}
			}
		}
	}
	// printM(m)
	// fmt.Println()
	// printM(n)

	return m[rows-1][cols-1]
}
