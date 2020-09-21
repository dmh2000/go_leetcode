package lcjune

import "sort"

func largestDivisibleSubset(nums []int) []int {
	var r [][]int
	var i int
	var j int
	var k int
	var max int
	r = [][]int{}

	if len(nums) <= 1 {
		return nums
	}

	sort.IntSlice.Sort(nums)

	for i = 0; i < len(nums); i++ {
		r = append(r, []int{nums[i]})
	}

	max = len(r[0])
	k = 0
	for i = 0; i < len(nums); i++ {
		for j = 0; j < i; j++ {
			if nums[i]%nums[j] == 0 && len(r[i]) < len(r[j])+1 {
				copy(r[i], r[j])
				r[i] = append(r[i], nums[i])
				if len(r[i]) > max {
					max = len(r[i])
					k = i
				}
			}
		}
	}
	return r[k]
}
