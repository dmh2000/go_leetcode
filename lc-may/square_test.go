package lcmay

import (
	"fmt"
	"math"
	"testing"
)

func verifySquare(n int) bool {
	if n == 0 {
		return false
	}
	x := float64(n)
	y := math.Sqrt(x)
	m := int(y)
	if m*m == n {
		return true
	}
	return false
}

func TestSquare(t *testing.T) {
	var a bool
	var b bool

	for i := 9; i < 1000; i++ {
		a = verifySquare(i)
		b = isPerfectSquare(i)
		fmt.Println(i, a, b)
		if a != b {
			t.Error(b, " should be ", a)
		}
	}
}
