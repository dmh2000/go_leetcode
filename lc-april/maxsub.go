package lc

func maxSubArray(nums []int) int {
	var sum int
	var big int
	sum = nums[0]
	big = sum
	for i := 1; i < len(nums); i++ {
		n := nums[i]
		m := sum + n
		if n > 0 {
			// sum will go up
			if sum < 0 {
				// discard sum and start over
				sum = n
			} else {
				// increased, keep going
				sum = m
			}
		} else {
			// sum will go down
			if m < 0 {
				// discard and start over
				sum = n
			} else {
				// still positive, keep going
				sum = m
			}
		}
		if sum > big {
			big = sum
		}
	}
	return big
}
