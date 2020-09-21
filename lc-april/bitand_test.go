package lc

import (
	"fmt"
	"testing"
)

func TestBitand(t *testing.T) {
	var b int
	var c int
	for m := 1; m < 1000; m++ {
		fmt.Println(m)
		for n := m; n < 1000; m++ {
			b = rangeBitWiseAnd(m, n)
			c = rangeBitWiseAnd2(m, n)
			if b != c {
				t.Error(c, " should be ", b)
			}
		}
	}
}
