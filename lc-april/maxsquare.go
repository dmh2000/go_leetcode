package lc

import "fmt"


func printMatrix(m [][]byte) {
	rows := len(m)
	cols := len(m[0])
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
            fmt.Printf("%v ", string(m[r][c]))
		}
		fmt.Println()
	}

}

func sqmax(m int, n int) int {
	if m > n {
		return m
	}
	return n

}

func sqmin(m int, n int) int {
	if m < n {
		return m
	}
	return n

}

func msquare(m [][]byte, r int, c int) int {
	rows := len(m)
	cols := len(m[0])
	x := 0
	y := 0
	minx := 2147483647
	for i := r; i < rows && m[i][c] != '0'; i++ {
		y++
		x = 0
		for j := c; j < cols && m[i][j] != '0'; j++ {
			x++
		}
		minx = sqmin(minx, x)
		if y >= minx {
			break
		}
	}
	minsq := sqmin(x, y)

	return minsq
}

func maximalSquare(matrix [][]byte) int {
	// printMatrix(matrix)

	rows := len(matrix)
    if rows == 0 {
        return 0
    }
	cols := len(matrix[0])
    if cols == 0 {
        return 0
    }
    
    // square of 1
    if rows == 1 {
        for _,v := range matrix[0] {
        if v == '1' {
            return 1
        }
        }
        return 0
    }
    
    if cols == 1 {
        for _,v := range matrix {
            if v[0] == '1' {
                return 1
            }
        }
        return 0
    }
    
	maxsq := 0
	sq := 0
    
	for r := 0; r < rows; {
		for c := 0; c < cols; c++ {
			if matrix[r][c] == '0' {
				continue
			}
			if matrix[r][c] == '1' {
				sq = msquare(matrix, r, c)
				maxsq = sqmax(maxsq, sq)
			}
			if sq == 0 {
				c++
			} else {
				c += sq
			}
		}
        if sq == 0 {
            r++
        } else {
            r += sq
        }
	}

	//fmt.Println()
	//printMatrix(matrix)
	return maxsq * maxsq
}
