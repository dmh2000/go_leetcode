package lc 

import "fmt"


// after an hour I had to research this one but I didn't 
// copy the code, just the gist of how it would work
func longestCommonSubsequence(text1 string, text2 string) int {

	rows := len(text1) + 1 // test1
	cols := len(text2) + 1 // text2
	m := make([][]int, rows)

    // i tried using an array of [1000][1000] instead of a slice but
    // the runtime was similar but used more memory
	// all cells default to 0 so the
    // first row and first columns cells are 
    // how they need to be initialized
	for i := 0; i < rows; i++ {
    	m[i] = make([]int, cols)
	}
	// first row remain all zeros
	// first column remains all zeros
	// that's why the loops start at 1
	// each character in the row string
	for r := 1; r < rows; r++ {
		// for each character in the column string
		for c := 1; c < cols; c++ {
			// do the row/col characters match?
			if text1[r-1] == text2[c-1] {
				// characters match, get value from left-upper diagonal
				// add 1 for the match
				m[r][c] = m[r-1][c-1] + 1
			} else {
				// carry forward max of left element and upper element
				if m[r-1][c] > m[r][c-1] {
					m[r][c] = m[r-1][c]
				} else {
					m[r][c] = m[r][c-1]
				}
			}
		}
	}

	fmt.Printf("    ")
	for c:=0;c<len(text2);c++ {
		fmt.Printf("%v ", string(text2[c]))
	}
	fmt.Println()
			
	for r:=0;r<rows;r++ {
		if r > 0 {
			fmt.Printf("%v ",string(text1[r-1]))
		} else {
			fmt.Printf("  ")
		}
		for c:=0;c<cols;c++ {
			fmt.Printf("%v ", m[r][c])
		}
		fmt.Println()
	}

	return m[rows-1][cols-1]
}