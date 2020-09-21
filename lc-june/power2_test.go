package lcjune

import (
	"fmt"
	"testing"
)

func TestPower2(t *testing.T) {

	for n := -2147483648; n < 0; n++ {
		if isPowerOfTwo(n) {
			fmt.Printf("%08x\n", n)
		}
	}
	fmt.Println()
}
