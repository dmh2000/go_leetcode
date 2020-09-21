package lcjune

import (
	"fmt"
	"testing"
)

func TestRecon1(t *testing.T) {
	var a [][]int
	// var b [][]int
	a = [][]int{
		{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2},
	}

	// b = [][]int{
	// 	{5, 0}, {7, 0}, {5, 2}, {6, 1}, {4, 4}, {7, 1},
	// }
	a = reconstructQueue(a)
	fmt.Println(a)
}
