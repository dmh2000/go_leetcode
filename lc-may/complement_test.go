package lcmay

import (
	"fmt"
	"math"
	"testing"
)

func TestComplement(t *testing.T) {
	for i := 0; i <= 31; i++ {
		findComplement(i)
	}
}

func TestLog2(t *testing.T) {
	for i := 1; i < 32; i++ {
		n := int(math.Log2(float64(i)))
		m := ilog2(i)
		fmt.Printf("%08b %v %v\n", i, n, m)
	}
}
