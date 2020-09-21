package lc

import (
	"fmt"
	"testing"
)

func TestLastStone(t *testing.T) {
	a := []int{2, 7, 4, 1, 8, 1}
	w := lastStoneWeight(a)
	fmt.Println(w)
}
