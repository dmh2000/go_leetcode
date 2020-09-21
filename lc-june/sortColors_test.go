package lcjune

import (
	"fmt"
	"testing"
)

func intSliceEqual(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true

}

func TestSortColors1(t *testing.T) {
	var a []int
	var b []int
	a = []int{2, 0, 2, 1, 1, 0}
	b = []int{2, 0, 2, 1, 1, 0}
	sortColors(a)
	countColors(b)
	if !intSliceEqual(a, b) {
		t.Error(a, "should be", b)
	}
	fmt.Println(a)
	fmt.Println(b)

	a = []int{1, 0, 1, 0, 1, 0}
	b = []int{1, 0, 1, 0, 1, 0}
	sortColors(a)
	countColors(b)
	if !intSliceEqual(a, b) {
		t.Error(a, "should be", b)
	}
	fmt.Println(a)
	fmt.Println(b)
}

func TestSortColors2(t *testing.T) {
	var a []int
	var b []int
	a = []int{2, 2, 2, 2, 2, 2}
	b = []int{2, 2, 2, 2, 2, 2}
	sortColors(a)
	countColors(b)
	if !intSliceEqual(a, b) {
		t.Error(a, "should be", b)
	}
	fmt.Println(a)
	fmt.Println(b)

	a = []int{0, 0, 0, 0, 0, 0}
	b = []int{0, 0, 0, 0, 0, 0}
	sortColors(a)
	countColors(b)
	if !intSliceEqual(a, b) {
		t.Error(a, "should be", b)
	}
	fmt.Println(a)
	fmt.Println(b)
}

func TestSortColors3(t *testing.T) {
	var a []int
	var b []int
	a = []int{2, 2}
	b = []int{2, 2}
	copy(b, a)
	sortColors(a)
	countColors(b)
	if !intSliceEqual(a, b) {
		t.Error(a, "should be", b)
	}
	fmt.Println(a)
	fmt.Println(b)

	a = []int{0, 0}
	b = []int{0, 0}
	sortColors(a)
	countColors(b)
	if !intSliceEqual(a, b) {
		t.Error(a, "should be", b)
	}
	fmt.Println(a)
	fmt.Println(b)

	a = []int{1, 1}
	b = []int{1, 1}
	sortColors(a)
	countColors(b)
	if !intSliceEqual(a, b) {
		t.Error(a, "should be", b)
	}
	fmt.Println(a)
	fmt.Println(b)

	a = []int{0}
	b = []int{0}
	sortColors(a)
	countColors(b)
	if !intSliceEqual(a, b) {
		t.Error(a, "should be", b)
	}
	fmt.Println(a)
	fmt.Println(b)

	a = []int{1}
	b = []int{1}
	sortColors(a)
	countColors(b)
	if !intSliceEqual(a, b) {
		t.Error(a, "should be", b)
	}
	fmt.Println(a)
	fmt.Println(b)

	a = []int{2}
	b = []int{2}
	sortColors(a)
	countColors(b)
	if !intSliceEqual(a, b) {
		t.Error(a, "should be", b)
	}
	fmt.Println(a)
	fmt.Println(b)
}

func TestSortColors4(t *testing.T) {
	var a []int
	var b []int
	a = []int{1, 2, 0}
	b = []int{1, 2, 0}
	copy(b, a)
	sortColors(a)
	countColors(b)
	fmt.Println(a)
	fmt.Println(b)
}
