package lc

import (
	"testing"
)

func testCases() [][]int {
	var a [][]int

	// create the array of cases
	a = make([][]int, 0)

	// add test cases
	a = append(a, []int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	a = append(a, []int{7, 8, 6, 2, 5, 1, 8, 3, 7})
	a = append(a, []int{9, 7, 1, 1, 1, 1, 1, 1, 1})
	a = append(a, []int{9, 9, 1, 1, 1, 1, 1, 1, 1})

	return a
}

var answers = []int{
	49, 56, 8, 9,
}

// TestMaxArea ...
func TestMaxArea1(t *testing.T) {
	var cases [][]int

	cases = testCases()
	for i := 0; i < 1000000; i++ {
		for i, v := range cases {
			area := maxAreaAllPairs1(v)
			if area != answers[i] {
				t.Error("should be ", answers[i], " got ", area)
			}

		}
	}
}

func TestMaxArea2(t *testing.T) {
	var cases [][]int

	cases = testCases()
	for i := 0; i < 1000000; i++ {
		for i, v := range cases {
			area := maxAreaAllPairs2(v)
			if area != answers[i] {
				t.Error("should be ", answers[i], " got ", area)
			}
		}
	}
}
