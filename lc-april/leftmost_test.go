package lc

import (
	"fmt"
	"testing"
)

func TestLeftmost(t *testing.T) {
	a := [][]int{
		{0, 0, 0, 1},
		{0, 0, 1, 1},
		{0, 1, 1, 1},
	}

	b := leftMostColumnWithOne(a)
	fmt.Println("result", b)

	a = [][]int{
		{1, 1, 1, 1, 1},
		{0, 0, 0, 1, 1},
		{0, 0, 1, 1, 1},
		{0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0}}
	b = leftMostColumnWithOne(a)
	fmt.Println("result", b)
}
