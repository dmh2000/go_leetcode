package lc

import (
	"fmt"
	"testing"
)

func cases() [][]int {
	var a [][]int

	a = make([][]int, 0)
	a = append(a, []int{0, 1, 2, 3, 4, 5})
	a = append(a, []int{0, 1, 0, 3, 4, 5})
	a = append(a, []int{0, 0, 1})
	a = append(a, []int{1, 1, 1, 0, 0})

	return a
}

func TestZeros(t *testing.T) {
	for _, v := range cases() {
		fmt.Println("------------------")
		fmt.Println(v)
		moveZeros(v)
		fmt.Println(v)
	}
}
