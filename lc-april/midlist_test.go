package lc

import (
	"fmt"
	"testing"
)

func midList(a []int) int {
	i := 0
	j := 0
	for {
		if j >= len(a) {
			i++
			break
		}
		j++
		if j >= len(a) {
			i++
			break
		}
		j++
		i++
	}
	return i
}
func TestMidList(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}

	n := midList(a)
	fmt.Println(n)
}
