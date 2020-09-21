package lc

func minP(r int, c int, grid [][]int, sum int, minsum *int) {
	rows := len(grid)
	cols := len(grid[0])

	if sum > *minsum {
		return
	}

	if r == rows-1 && c == cols-1 {
		sum += grid[r][c]
		if sum < *minsum {
			*minsum = sum
		}
		return
	}

	// add in the sum
	sum += grid[r][c]

	// move right
	if c < cols-1 {
		minP(r, c+1, grid, sum, minsum)
	}

	// move down
	if r < rows-1 {
		minP(r+1, c, grid, sum, minsum)
	}
}

// recursive depth first search version
func minPathSum1(grid [][]int) int {
	sum := 0
	minsum := 2147483647

	if len(grid) == 0 {
		return 0
	}
	if len(grid[0]) == 0 {
		return 0
	}

	minP(0, 0, grid, sum, &minsum)

	return minsum
}

func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	if len(grid[0]) == 0 {
		return 0
	}

	var rows = len(grid)
	var cols = len(grid[0])
	var r int
	var c int

	// make an aux array for the solution
	a := make([][]int, rows)
	for r = 0; r < rows; r++ {
		a[r] = make([]int, cols)
	}

	// compute sums across top row
	r = 0
	a[0][0] = grid[0][0]
	for c = 1; c < cols; c++ {
		a[r][c] = grid[r][c] + a[r][c-1]
	}
	// compute sums across left column
	c = 0
	for r = 1; r < rows; r++ {
		a[r][c] = grid[r][c] + a[r-1][c]
	}
	// compute sum across last column
	c = cols - 1
	for r = 1; r < rows; r++ {
		a[r][c] = grid[r][c] + a[r-1][c]
	}

	// compute remaining sums
	for r = 1; r < rows; r++ {
		for c = 1; c < cols; c++ {
			x := grid[r][c] + a[r-1][c]
			y := grid[r][c] + a[r][c-1]
			if x < y {
				a[r][c] = x
			} else {
				a[r][c] = y
			}
			// if current sum is greater than existing final sum, quit this row
			if a[r][c] > a[rows-1][cols-1] {
				break
			}
		}
	}
	return a[rows-1][cols-1]
}
