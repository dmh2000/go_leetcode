package lcjune


import (
	"fmt"
	"sort"
	"testing"
)

func printSlice1D(a []int) {
	for i := 0; i < len(a); i++ {
		fmt.Printf("%2v,", a[i])
	}
	fmt.Println()
}

func printSlice2D(a [][]int) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			fmt.Printf("%2v,", a[i][j])
		}
		fmt.Println()
	}
}

func TestDivSub1(t *testing.T) {
	var a []int
	var b []int

	a = []int{2}
	b = []int{2}
	a = largestDivisibleSubset(a)
	sort.IntSlice.Sort(a)
	if !intSliceEqual(a, b) {
		fmt.Println(a)
		t.Error(a, "should be", b)
	}
	fmt.Println(a)

	a = []int{2, 3, 7}
	b = []int{2}
	a = largestDivisibleSubset(a)
	sort.IntSlice.Sort(a)
	if !intSliceEqual(a, b) {
		fmt.Println(a)
		t.Error(a, "should be", b)
	}
	fmt.Println(a)

}

func TestDivSub2(t *testing.T) {
	var a []int
	var b []int

	a = []int{1, 2, 4, 8, 16}
	b = []int{1, 2, 4, 8, 16}
	a = largestDivisibleSubset(a)
	sort.IntSlice.Sort(a)
	if !intSliceEqual(a, b) {
		fmt.Println(a)
		t.Error(a, "should be", b)
	}

	a = []int{556, 669}
	a = largestDivisibleSubset(a)
	sort.IntSlice.Sort(a)
	fmt.Println(a)

	a = []int{10, 240, 4, 8}
	b = []int{4, 8, 240}
	a = largestDivisibleSubset(a)
	sort.IntSlice.Sort(a)
	if !intSliceEqual(a, b) {
		fmt.Println(a)
		t.Error(a, "should be", b)
	}

	a = []int{3, 4, 16, 8}
	b = []int{4, 8, 16}
	a = largestDivisibleSubset(a)
	sort.IntSlice.Sort(a)
	if !intSliceEqual(a, b) {
		fmt.Println(a)
		t.Error(a, "should be", b)
	}

	a = []int{2, 3, 4, 8}
	b = []int{2, 4, 8}
	a = largestDivisibleSubset(a)
	sort.IntSlice.Sort(a)
	if !intSliceEqual(a, b) {
		fmt.Println(a)
		t.Error(a, "should be", b)
	}
}
