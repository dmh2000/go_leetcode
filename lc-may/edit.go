package lcmay

import "fmt"

func freqCount(b []byte) map[byte]int {
	var f map[byte]int

	f = make(map[byte]int)
	for _, v := range b {
		f[v]++
	}
	return f
}

func printB(b [][]byte) {
	for r := 0; r < len(b); r++ {
		fmt.Println(b[r])
	}
	fmt.Println()
}

func minDistance(word1 string, word2 string) int {
	var b1 []byte
	var b2 []byte
	var a [][]byte
	var r int
	var c int
	var rows int
	var cols int
	var min byte
	if word1 == "" {
		return len(word2)
	} else if word2 == "" {
		return len(word1)
	}

	b1 = []byte(word1)
	b2 = []byte(word2)

	// extend rows/cols by 1 to pad left column, upper row
	rows = len(b1) + 1
	cols = len(b2) + 1
	// set up the matrix
	a = make([][]byte, rows)
	for r = 0; r < rows; r++ {
		a[r] = make([]byte, cols)
	}

	// set [0] of each row
	// from 0 to length of w1
	for r := 0; r < rows; r++ {
		a[r][0] = byte(r)
	}
	// set [0] of each column
	// from 0 to length(w2)
	for c = 0; c < cols; c++ {
		a[0][c] = byte(c)
	}

	// start at [1][1] and process each
	// row,column
	for r = 1; r < rows; r++ {
		for c = 1; c < cols; c++ {
			if b1[r-1] != b2[c-1] {
				min = byte(255)
				// not-equal
				// get minimum of left, upper or diagnonal
				if a[r-1][c-1] < min {
					// diagonal : substitution
					min = a[r-1][c-1]
				}
				if a[r-1][c] < min {
					// upper : deletion
					min = a[r-1][c]
				}
				if a[r][c-1] < min {
					// left : insertion
					min = a[r][c-1]
				}
				a[r][c] = min + 1
			} else {
				// equal, get value from diagonal
				a[r][c] = a[r-1][c-1]
			}
		}
	}
	// printB(a)

	return int(a[rows-1][cols-1])
}
