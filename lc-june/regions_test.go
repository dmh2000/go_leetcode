package lcjune

import (
	"fmt"
	"testing"
)

func TestRegions1(t *testing.T) {
	var b [][]byte

	b = [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	}

	solve(b)
	fmt.Println()

	b = [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'O', 'X'},
	}
	solve(b)
	fmt.Println()

	b = [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
	}
	solve(b)
	fmt.Println()
}
