package lcmay

func majorityElement(nums []int) int {

	target := len(nums) / 2

	m := make(map[int]int)

	// frequency count
	for _, v := range nums {
		m[v]++
	}

	// find majority element
	r := 0
	for k, v := range m {
		if v > target {
			r = k
			break
		}
	}

	return r
}
