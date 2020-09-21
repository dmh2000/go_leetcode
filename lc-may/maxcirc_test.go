package lcmay

import (
	"fmt"
	"testing"
)

func TestMaxCirc1(t *testing.T) {
	var sum int

	sum = maxSubarraySumCircular([]int{-5, -2, 5, 6, -2, -7, 0, 2, 8})
	if sum != 14 {
		t.Error(sum, " should be ", 14)
		fmt.Println("---")
	}

	sum = maxSubarraySumCircular([]int{2, -2, 2, 7, 8, 0})
	if sum != 19 {
		t.Error(sum, " should be ", 19)
		fmt.Println("---")
	}

	sum = maxSubarraySumCircular([]int{2, -2, 2, 7, 8, -1})
	if sum != 18 {
		t.Error(sum, " should be ", 18)
		fmt.Println("---")
	}

	sum = maxSubarraySumCircular([]int{0, 5, 8, -9, 9, -7, 3, -2})
	if sum != 16 {
		t.Error(sum, " should be ", 16)
		fmt.Println("---")
	}

	sum = maxSubarraySumCircular([]int{5, -3, 5})
	if sum != 10 {
		t.Error(sum, " should be ", 10)
		fmt.Println("---")
	}

	sum = maxSubarraySumCircular([]int{-7, 1, 2, 7, -2, -5})
	if sum != 10 {
		t.Error(sum, " should be ", 10)
		fmt.Println("---")
	}

	sum = maxSubarraySumCircular([]int{1, -2, 3, -2})
	if sum != 3 {
		t.Error(sum, " should be ", 3)
		fmt.Println("---")
	}

	sum = maxSubarraySumCircular([]int{3, -1, 2, -1})
	if sum != 4 {
		t.Error(sum, " should be ", 4)
		fmt.Println("---")
	}

	sum = maxSubarraySumCircular([]int{3, -2, 2, -3})
	if sum != 3 {
		t.Error(sum, " should be ", 3)
		fmt.Println("---")
	}

	sum = maxSubarraySumCircular([]int{-3, -2, -1})
	if sum != -1 {
		t.Error(sum, " should be ", -1)
		fmt.Println("---")
	}

	sum = maxSubarraySumCircular([]int{5, 5, 5})
	if sum != 15 {
		t.Error(sum, " should be ", 15)
		fmt.Println("---")
	}

	sum = maxSubarraySumCircular([]int{0, 0, 0, 0, 0, 0})
	if sum != 0 {
		t.Error(sum, " should be ", 0)
		fmt.Println("---")
	}
}
