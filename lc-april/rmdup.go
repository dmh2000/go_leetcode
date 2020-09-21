package lc

import "fmt"

func removeDuplicates(nums []int) int {
	var i int
	var j int
	var count int
	i = 0
	j = 1
	count = 0
	for j < len(nums) {
		if nums[i] == nums[j] {
			j++
			count++
		} else {
			i++
			nums[i] = nums[j]
			j++
		}

		fmt.Println(nums)
	}
	return len(nums) - count
}
