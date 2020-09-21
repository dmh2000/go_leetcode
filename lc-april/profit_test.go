package lc

import (
	"fmt"
	"testing"
)

func profitCases() [][]int {
	return [][]int{
		{7, 1, 5, 3, 6, 4},
		{1, 2, 3, 4, 5},
		{7, 5, 4, 3, 1},
	}
}

func TestProfit(t *testing.T) {
	for _, a := range profitCases() {
		p := maxProfit(a)
		fmt.Println(p)
	}
}
