package lcmay

import (
	"fmt"
	"testing"
)

func TestSpan1(t *testing.T) {
	var span StockSpanner
	var n int
	var inp []int
	var out []int

	inp = []int{100, 80, 60, 70, 60, 75, 85}
	out = []int{1, 1, 1, 2, 1, 4, 6}

	span = SpanConstructor()

	for i := 0; i < len(inp); i++ {
		n = span.Next(inp[i])
		fmt.Printf("%3v %3v %3v\n", inp[i], out[i], n)
		if n != out[i] {
			t.Error(n, " should be ", out[i])
		}
	}
	fmt.Println()

	// [[],[28],[14],[28],[35],[46],[53],[66],[80],[87],[88]]
	inp = []int{28, 14, 28, 35, 46, 53, 66, 80, 87, 88}
	out = []int{1, 1, 3, 4, 5, 6, 7, 8, 9, 10}

	span = SpanConstructor()

	for i := 0; i < len(inp); i++ {
		n = span.Next(inp[i])
		fmt.Printf("%3v %3v %3v\n", inp[i], out[i], n)
		if n != out[i] {
			t.Error(n, " should be ", out[i])
		}
	}
	fmt.Println()
}
