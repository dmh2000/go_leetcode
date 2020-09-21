package lc

import (
	"fmt"
	"testing"
)

func TestDigits(t *testing.T) {
	var d []int
	d = digits(0)
	fmt.Println(d)
	j := 2
	for i := 0; i < 10; i++ {
		d = digits(j)
		fmt.Println(j, d)
		j = j*10 + 2
	}
}

func TestSquare(t *testing.T) {
	for i := 0; i < 10; i++ {
		d := digits(i)
		s := square(d)
		fmt.Println(s)
	}
}

func TestHappy(t *testing.T) {
	a := []int{
		19,
		23,
		28,
		31,
		32,
		44,
		49,
		68,
		70,
		79,
		82,
		86,
		91,
		94,
		97,
	}
	m := make(map[int]bool)
	for _, v := range a {
		m[v] = true
	}

	for n := 15; n < 100; n++ {
		h := isHappy(n)
		if h != m[n] {
			t.Errorf("%v should be %v", n, m[n])
		}
	}
}
