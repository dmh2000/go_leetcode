package lc

import (
	"fmt"
	"testing"
)

func checkPossibility(nums []int) bool {

	increasing := func(a []int) bool {
		for i := 0; i < len(a)-1; i++ {
			if a[i] > a[i+1] {
				return false
			}
		}
		return true
	}

	// replace each number with the preivous
	// and if it fixes it then its done
	for i := range nums {
		ni := nums[i]
		if i > 0 {
			nums[i] = nums[i-1]
		} else {
			nums[i] = -2147483648
		}
		if increasing(nums) {
			return true
		}
		nums[i] = ni
	}

	return false
}

func TestNonDec(t *testing.T) {
	a := []int{3, 4, 2, 3}
	b := checkPossibility(a)
	if b {
		t.Error(b, " should be ", false)
	}
	fmt.Println(a)

	a = []int{4, 2, 3}
	b = checkPossibility(a)
	if !b {
		t.Error(b, " should be ", true)
	}
	fmt.Println(a)

	a = []int{1, 3, 2}
	b = checkPossibility(a)
	if !b {
		t.Error(b, " should be ", true)
	}
	fmt.Println(a)

	a = []int{4, 2, 1}
	b = checkPossibility(a)
	if b {
		t.Error(b, " should be ", false)
	}
	fmt.Println(a)

	a = []int{2, 3, 3, 2, 4}
	b = checkPossibility(a)
	if !b {
		t.Error(b, " should be ", true)
	}
	fmt.Println(a)
}
