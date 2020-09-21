package lcjune

func countColors(nums []int) {
	var m map[int]int

	m = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}

	k := 0
	for i := 0; i <= 2; i++ {
		for j := 0; j < m[i]; j++ {
			nums[k] = i
			k++
		}
	}
}

func sortColors1(nums []int) {
	var i int
	var j int
	var k int

	if len(nums) <= 1 {
		return
	}

	j = len(nums) - 1
	k = 0
loop:
	for i < len(nums) && k < j {
		if nums[i] == 2 && i < j {
			// swap to back
			nums[i], nums[j] = nums[j], nums[i]
			for nums[j] == 2 {
				j--
				if j < 0 {
					break loop
				}
			}
		} else if nums[i] == 0 && i > k {
			// swap to front
			nums[i], nums[k] = nums[k], nums[i]
			for nums[k] == 0 {
				k++
				if k == len(nums) {
					break loop
				}
			}
		} else {
			i++
		}
	}
}

func sortColors(nums []int) {
	var i int
	var j int
	var k int

	if len(nums) <= 1 {
		return
	}

	j = len(nums) - 1
	k = 0
	for i <= j {
		if nums[i] == 2 {
			// swap to back
			nums[i], nums[j] = nums[j], nums[i]
			j--
		} else if nums[i] == 0 {
			// swap to front
			nums[i], nums[k] = nums[k], nums[i]
			k++
			i++
		} else {
			// no swap, next i
			i++
		}
	}
}
