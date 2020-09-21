package lc

import "testing"

func jtcases() ([][]int, []bool) {
	var j [][]int
	var b []bool

	j = append(j, []int{5, 4, 3, 2, 1, 0, 0, 8})
	b = append(b, false)

	j = append(j, []int{2, 3, 1, 1, 4})
	b = append(b, true)

	j = append(j, []int{3, 2, 1, 0, 4})
	b = append(b, false)

	j = append(j, []int{3, 3, 0, 0, 4})
	b = append(b, true)

	j = append(j, []int{1, 3, 0, 0, 4})
	b = append(b, true)
	return j, b
}

func TestJump1(t *testing.T) {

	j, b := jtcases()

	for i, v := range j {
		k := canJump(v)
		if k != b[i] {
			t.Error(k, " should be ", b[i])
		}
	}
}
