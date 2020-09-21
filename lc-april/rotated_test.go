package lc

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {

	var a []int
	var m int
	a = []int{0, 1, 2, 3, 4, 5}
	for i := 0; i < len(a); i++ {
		m = binarySearch(a, i)
		if m != i {
			t.Error(m, " should be ", i)
		}
	}
	a = []int{0, 1, 2, 3, 4, 5, 6}
	for i := 0; i < len(a); i++ {
		m = binarySearch(a, i)
		if m != i {
			t.Error(m, " should be ", i)
		}
	}
}

func TestRotated(t *testing.T) {

	var a []int
	a = []int{0, 1, 2, 3, 4, 5}
	for i := len(a) - 1; i >= 0; i-- {
		// rotate the slice
		a = append(a[1:], a[0:1]...)
		fmt.Println(a)
		offset := findOffset2(0, len(a)-1, a)
		if offset != i {
			t.Error(offset, " should be ", i)
		}
	}
	a = []int{5, 6, 7, 8, 9, 1, 2, 3, 4}
	offset := findOffset2(0, len(a)-1, a)
	fmt.Println(offset)

}

func TestRotated2(t *testing.T) {
	var a []int
	var m int
	var offset int

	offset = 0
	a = []int{0, 1, 2, 3, 4, 5}
	for i := 0; i < len(a); i++ {
		m = search(a, i)
		if m != (i+offset)%len(a) {
			t.Error(m, " should be ", (i+offset)%len(a))
		}
	}

	offset = 0
	a = []int{0, 1, 2, 3, 4, 5, 6}
	for i := 0; i < len(a); i++ {
		m = search(a, i)
		if m != (i+offset)%len(a) {
			t.Error(m, " should be ", (i+offset)%len(a))
		}
	}

	offset = 5
	a = []int{2, 3, 4, 5, 6, 0, 1}
	for i := 0; i < len(a); i++ {
		m = search(a, i)
		if m != (i+offset)%len(a) {
			t.Error(m, " should be ", (i+offset)%len(a))
		}
	}

	offset = 2
	a = []int{6, 7, 0, 1, 2, 3, 4, 5}
	for i := 0; i < len(a); i++ {
		m = search(a, i)
		if m != (i+offset)%len(a) {
			t.Error(m, " should be ", (i+offset)%len(a))
		}
	}

	offset = 8
	a = []int{1, 2, 3, 4, 5, 6, 7, 8, 0}
	for i := 0; i < len(a); i++ {
		m = search(a, i)
		if m != (i+offset)%len(a) {
			t.Error(m, " should be ", (i+offset)%len(a))
		}
	}
}
