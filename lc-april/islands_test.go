package lc

import (
	"testing"
)

func TestIslands1(t *testing.T) {
	grid := make([][]byte, 0)
	grid = append(grid, []byte{1, 1, 1, 1, 0})
	grid = append(grid, []byte{1, 1, 0, 1, 0})
	grid = append(grid, []byte{1, 1, 0, 0, 0})
	grid = append(grid, []byte{0, 0, 0, 0, 0})

	count := numIslands(grid)
	if count != 1 {
		t.Error(count, " should be ", 1)
	}

	grid = make([][]byte, 0)
	grid = append(grid, []byte{1, 1, 0, 0, 0})
	grid = append(grid, []byte{1, 1, 0, 0, 0})
	grid = append(grid, []byte{0, 0, 1, 0, 0})
	grid = append(grid, []byte{0, 0, 0, 1, 1})

	count = numIslands(grid)
	if count != 3 {
		t.Error(count, " should be ", 3)
	}

}
