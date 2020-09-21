package lcmay

import (
	"fmt"
	"testing"
)

func printImage(a [][]int) {
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
}
func TestFlood1(t *testing.T) {
	var a [][]int

	a = make([][]int, 8)
	a[0] = []int{1, 1, 2, 2, 2, 2, 2, 2}
	a[1] = []int{1, 1, 1, 1, 1, 1, 1, 2}
	a[2] = []int{1, 1, 1, 2, 1, 1, 1, 2}
	a[3] = []int{1, 1, 1, 1, 1, 1, 1, 2}
	a[4] = []int{1, 1, 1, 1, 1, 1, 1, 2}
	a[5] = []int{1, 1, 1, 1, 1, 1, 1, 1}
	a[6] = []int{1, 1, 1, 1, 1, 1, 1, 1}
	a[7] = []int{1, 1, 1, 1, 1, 1, 1, 1}

	printImage(a)

	a = floodFill(a, 0, 0, 4)

	printImage(a)
}

func TestFlood2(t *testing.T) {
	var a [][]int

	a = make([][]int, 2)
	a[0] = []int{0, 0, 0}
	a[1] = []int{0, 1, 1}

	printImage(a)

	a = floodFill(a, 1, 1, 1)

	printImage(a)
}
