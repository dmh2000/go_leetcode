package lc

func moveZeros(nums []int) {
	count := 0
	i := 0
	// the catch here is len(nums) changes when 0 is removed
	for i < len(nums) {
		if nums[i] == 0 {
			// count it
			count++
			// remove it
			nums = append(nums[0:i], nums[i+1:]...)
			// don't increment i
		} else {
			i++
		}
	}
	// append zeros
	for i := 0; i < count; i++ {
		nums = append(nums, 0)
	}
}
