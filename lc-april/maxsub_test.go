package lc

import "testing"

// TestMaxSum ...
func TestMaxSum(t *testing.T) {
	nums := [9]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	n := maxSubArray(nums[:])
	if n != 6 {
		t.Error("n should be 6, go t", n)
	}

	nums = [9]int{-2, 1, -3, 4, -5, 2, 1, -5, 4}
	n = maxSubArray(nums[:])
	if n != 4 {
		t.Error("n should be 4, got ", n)
	}

	nums = [9]int{-2, -1, -3, -4, -5, -2, -1, -5, -4}
	n = maxSubArray(nums[:])
	if n != -1 {
		t.Error("n should be -1, got ", n)
	}

	nums = [9]int{-2, -1, -3, 4, -1, -1, -1, -5, 3}
	n = maxSubArray(nums[:])
	if n != 4 {
		t.Error("n should be 4, got ", n)
	}

	nums = [9]int{-2, -1, -3, 4, 1, 1, -1, -5, 3}
	n = maxSubArray(nums[:])
	if n != 6 {
		t.Error("n should be 6, got ", n)
	}
}
