package lc

func productExceptSelf1(nums []int) []int {

	prefix := make([]int, len(nums))
	suffix := make([]int, len(nums))

	prefix[0] = 1
	for i := 1; i < len(nums); i++ {
		prefix[i] = nums[i-1] * prefix[i-1]
	}

	suffix[len(nums)-1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		suffix[i] = nums[i+1] * suffix[i+1]
	}

	r := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		r[i] = prefix[i] * suffix[i]
	}

	return r
}

func productExceptSelf(nums []int) []int {
	n := len(nums)
	r := make([]int, len(nums))

	r[0] = 1
	for i := 1; i < n; i++ {
		r[i] = r[i-1] * nums[i-1]
	}

	k := nums[n-1]
	for i := n - 2; i >= 0; i-- {
		r[i] = k * r[i]
		k = k * nums[i]
	}
	return r
}
