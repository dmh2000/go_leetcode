package lcmay

import (
	"strings"
)

func removeKdigitsBruteForce(num string, k int) string {
	var min string
	var t string

	if len(num) == k {
		return "0"
	}

	if k == 0 {
		return num
	}

	min = num
	t = num
	for i := 0; i < k; i++ {
		// remove digits until smallest one found
		for j := 0; j < len(t); j++ {
			// remove digit[j]
			s := t[0:j] + t[j+1:]

			// keep min value
			b := strings.Compare(s, min)
			if b < 0 {
				min = s
			}
			//fmt.Println(min, s)
		}
		t = min
		//fmt.Println()
	}

	// remove leading zeros
	for len(min) > 0 && min[0] == '0' {
		min = min[1:]
	}

	if len(min) == 0 {
		min = "0"
	}

	return min
}

func removeK(num string, k int, s *string) {
	if k == 0 {
		*s += num
		return
	}

	if len(num) <= k {
		return
	}

	idx := 0
	min := byte('9')
	for i := 0; i <= k; i++ {
		if num[i] < min {
			min = num[i]
			idx = i
		}
	}
	*s = *s + string(min)

	removeK(num[idx+1:], k-idx, s)
}

func removeKdigits(num string, k int) string {
	var t string
	if len(num) == k {
		return "0"
	}

	if k == 0 {
		return num
	}

	removeK(num, k, &t)

	// remove leading zeros
	for len(t) > 0 && t[0] == '0' {
		t = t[1:]
	}

	if len(t) == 0 {
		t = "0"
	}

	return t
}
