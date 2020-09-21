package lcmay

import (
	"fmt"
	"testing"
)

func TestIntersect(t *testing.T) {
	a := [][]int{
		{0, 2}, {5, 10}, {13, 23}, {24, 25},
	}
	b := [][]int{
		{1, 5}, {8, 12}, {15, 24}, {25, 26},
	}

	r := intervalIntersection(a, b)
	c := [][]int{
		{1, 2}, {5, 5}, {8, 10}, {15, 23}, {24, 24}, {25, 25},
	}
	fmt.Println(r)
	fmt.Println(c)
}

func TestIntersect2(t *testing.T) {
	a := [][]int{
		{0, 2}, {5, 10}, {13, 23}, {24, 25},
	}
	b := [][]int{
		{1, 5}, {8, 12}, {15, 24}, {25, 26}, {30, 31}, {40, 41},
	}

	r := intervalIntersection(a, b)
	c := [][]int{
		{1, 2}, {5, 5}, {8, 10}, {15, 23}, {24, 24}, {25, 25}, {30, 31}, {40, 41},
	}
	fmt.Println(r)
	fmt.Println(c)
}

func TestIntersect3(t *testing.T) {
	a := [][]int{
		{14, 16},
	}
	b := [][]int{
		{7, 13}, {16, 20},
	}

	r := intervalIntersection(a, b)
	c := [][]int{
		{16, 16},
	}
	fmt.Println(r)
	fmt.Println(c)
}
