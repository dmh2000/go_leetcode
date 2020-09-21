package lcmay

func searchSingle(a []int, left int, right int) (int, bool) {
	var n int
	var ok bool

	if left >= right {
		return 0, false
	}

	// get midpoint
	mid := left + ((right - left) / 2)

	// check it
	target := a[mid]

	if left < mid && mid < right {
		// middle
		if a[mid-1] != target && target != a[mid+1] {
			return target, true
		}
	} else if left == mid {
		// left end
		if target != a[mid+1] {
			return target, true
		}
	} else if right == mid {
		// right end
		if a[mid-1] != target {
			return target, true
		}
	}

	n, ok = searchSingle(a, left, mid-1)
	if ok {
		return n, ok
	}

	n, ok = searchSingle(a, mid, right+1)
	if ok {
		return n, ok
	}

	return 0, false
}

func singleNonDuplicate(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	// left end
	if nums[0] != nums[1] {
		return nums[0]
	}

	// right end
	if nums[len(nums)-2] != nums[len(nums)-1] {
		return nums[len(nums)-1]
	}

	n, _ := searchSingle(nums, 0, len(nums)-1)

	return n
}
