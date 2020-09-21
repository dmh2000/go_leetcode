package lc

// =================================================================
// I needed to look at the solutions to figure this one out
// =================================================================
func imax(a int, b int) int {
	if a > b {
		return a
	}
	return b

}

func findMaxLength(nums []int) int {
	m := make(map[int]int)

	m[0] = -1
	max := 0
	count := 0
	for i, v := range nums {
		if v == 1 {
			count++
		} else {
			count--
		}
		idx, ok := m[count]
		if ok {
			max = imax(max, i-idx)
		} else {
			m[count] = i
		}
	}

	return max
}
