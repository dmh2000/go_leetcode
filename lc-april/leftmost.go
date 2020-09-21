package lc

import "fmt"

/**
 * // This is the BinaryMatrix's API interface.
 * // You should not implement it, or speculate about its implementation
 */

func get(r int, c int, a [][]int) int {
	return a[r][c]
}

func dim(a [][]int) []int {
	rows := len(a)
	cols := len(a[0])
	return []int{rows, cols}
}

func binarySearchRow(cols int, row int, a [][]int) int {
	left := 0
	right := cols - 1
	mid := 0
	m := 1000

	for left <= right {
		mid = left + ((right - left) / 2)
		v := get(row, mid, a)
		if v == 1 {
			if mid < m {
				m = mid
			}
			// search left
			right = mid - 1
		} else {
			// search left
			left = mid + 1
		}
	}
	return m
}

func leftMostColumnWithOne(a [][]int) int {
	rows := len(a)
	cols := len(a[0])
	mincol := 1000
	for row := 0; row < rows; row++ {
		c := binarySearchRow(cols, row, a)
		fmt.Println(c)
		if c != -1 {
			if c < mincol {
				mincol = c
			}
		}
	}
	if mincol == 1000 {
		mincol = -1
	}
	return mincol
}
