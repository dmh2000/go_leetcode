package lc

func canJump(nums []int) bool {
	if len(nums) == 0 {
		return false
	} else if len(nums) == 1 {
		return true
	}

	// goal is last element in array
	goal := len(nums) - 1
	// start at 1 before the goal and go backwards
	for i := len(nums) - 2; i > 0; i-- {
		// break out the terms for clarity
		// get current jump count
		n := nums[i]
		// distance to goal for this step is from
		// total goal minus current index
		g := goal - i
		// if you can get from this position to the goal
		// then this position becomes the goal and you just
		// need to get to it to get to the final goal
		if n >= g {
			goal = i
		}
	}

	// if the count in the first element is >= the computed goal
	return nums[0] >= goal
}
