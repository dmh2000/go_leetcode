package lcmay

import "fmt"

var nn int

func setNN(n int) {
	nn = n
}

func isBadVersion(v int) bool {
	if v >= nn {
		return true
	}
	return false
}

func firstBadVersion(n int) int {
	left := 1
	right := n
	mid := 0

	fmt.Printf("first bad : %v : ", nn)
	for left <= right {
		mid = left + ((right - left) / 2)

		if isBadVersion(mid) {
			// if previous is not bad this is it
			if !isBadVersion(mid - 1) {
				fmt.Println(mid, nn)
				return mid
			}
			// go right
			right = mid - 1
		} else {
			// n is not bad
			// if next is bad, then that is it
			if isBadVersion(mid + 1) {
				fmt.Println(mid+1, nn)
				return mid + 1
			}
			left = mid + 1
		}

	}
	return -1

}
