package lc

import "fmt"

func binarySearch(a []int, target int) int {

	left := 0
	right := len(a) - 1
	mid := 0

	for left <= right {
		mid = left + ((right - left) / 2)
		if a[mid] == target {
			return mid
		}
		if target < a[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func findOffset(a []int) int {
	n := len(a)
	min := 2147483647
	offset := 0
	for i := 0; i < n; i++ {
		if a[i] < min {
			min = a[i]
			offset = i
		}
	}
	return offset
}

func findOffset2(lo int, hi int, a []int) int {
	if lo > hi {
		return -1
	}

	if lo == hi {
		return lo
	}

	// array is sorted
	if a[hi] > a[lo] {
		return lo
	}

	m := (lo + hi) / 2
	fmt.Println(a[lo], a[m], a[hi])
	if m == lo && a[m] > a[m+1] {
		return m + 1
	} else if a[m-1] > a[m] && a[m] < a[m+1] {
		return m
	} else if a[lo] <= a[m] {
		// lower range is sorted, look in upper range
		if (hi-m)%2 == 0 {
			return findOffset2(m+1, hi, a)
		}
		return findOffset2(m, hi, a)
	} else {
		// look in lower range
		return findOffset2(lo, m-1, a)
	}

}

func search(a []int, target int) int {
	n := len(a)

	offset := 0
	if a[0] > a[n-1] {
		// array is rotated
		offset = findOffset2(0, n-1, a)
	}

	x := findOffset(a)
	if offset != x {
		return -1
	}

	left := 0
	right := len(a) - 1
	mid := 0

	for left <= right {
		mid = left + ((right - left) / 2)
		m := (mid + offset) % n
		if a[m] == target {
			return m
		}
		if target < a[m] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}
