package lcjune

import "fmt"

func printBoardM(b [][]byte, m [][]int) {
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[i]); j++ {
			fmt.Printf("%v:%v ", string(b[i][j]), m[i][j])
		}
		fmt.Println()
	}
}

func printBoard(b [][]byte) {
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[i]); j++ {
			fmt.Printf("%v ", string(b[i][j]))
		}
		fmt.Println()
	}
}

func searchBoard(row int, col int, b [][]byte, m [][]int) {
	var rows = len(b)
	var cols = len(b[0])

	// row out of range?
	if row < 0 || row >= rows {
		return
	}

	// col out of range?
	if col < 0 || col >= cols {
		return
	}

	// not a O?
	if b[row][col] != 'O' {
		return
	}

	// already marked?
	if m[row][col] == 1 {
		return
	}

	// mark the cell
	m[row][col] = 1

	// search four directions
	// up
	searchBoard(row-1, col, b, m)
	// left
	searchBoard(row, col-1, b, m)
	// right
	searchBoard(row, col+1, b, m)
	// down
	searchBoard(row+1, col, b, m)

}

func solve(board [][]byte) {
	var marked [][]int
	var r int
	var c int

	// degenerate rows
	if len(board) <= 1 {
		return
	}

	// degenrate columns
	if len(board[0]) <= 1 {
		return
	}

	marked = make([][]int, len(board))
	for r := 0; r < len(board); r++ {
		marked[r] = make([]int, len(board[0]))
	}

	// traverse the edges and mark any zero's  on the edge
	// and any connected to zero's on the edge

	// top row
	r = 0
	for c = 0; c < len(board[r]); c++ {
		searchBoard(r, c, board, marked)
	}
	// left column
	c = 0
	for r = 0; r < len(board); r++ {
		searchBoard(r, c, board, marked)
	}
	// right colunn
	c = len(board[0]) - 1
	for r = 0; r < len(board); r++ {
		searchBoard(r, c, board, marked)
	}
	// bottom row
	r = len(board) - 1
	for c = 0; c < len(board[r]); c++ {
		searchBoard(r, c, board, marked)
	}

	// now flip any 'O' not marked
	for r = 0; r < len(board); r++ {
		for c = 0; c < len(board[0]); c++ {
			if board[r][c] == 'O' && marked[r][c] == 0 {
				board[r][c] = 'X'
			}
		}
	}
	printBoard(board)
}
