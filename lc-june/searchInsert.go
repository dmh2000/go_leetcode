package lcjune

func binarySearchInsert(a []int, target int) int {

	left := 0
	right := len(a) - 1
	mid := 0

	if len(a) == 0 {
		return -1
	}

	if len(a) == 1 {
		if target <= a[0] {
			return 0
		}
		return 1
	}

	for left <= right {
		mid = left + ((right - left) / 2)

		// target found
		if a[mid] == target {
			return mid
		} else if mid > 0 && mid < len(a)-1 {
			// somewhere in the middle
			if target < a[mid] {
				if a[mid-1] < target {
					return mid
				}
			}
		} else if mid == len(a)-1 {
			// at the end
			if target > a[mid] {
				return mid + 1
			} else if a[mid] >= target {
				return mid
			}
		} else if mid == 0 {
			// at the beginning
			if target < a[mid] {
				return 0
			}
		}

		if target < a[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func searchInsert(nums []int, target int) int {
	return binarySearchInsert(nums, target)
}
