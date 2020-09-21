package lc

func area(a int, b int, l int) int {
	// get least height
	h := a
	if b < a {
		h = b
	}

	return h * l
}

func maxAreaAllPairs1(height []int) int {
	a := 0
	pi := 0
	for i := 0; i < len(height); i++ {
		hi := height[i]
		if hi > pi {
			pi = hi
		} else {
			continue
		}
		pj := 0
		for j := len(height) - 1; j > i; j-- {
			hj := height[j]
			if hj > pj {
				pj = hj
			} else {
				continue
			}

			t := area(hi, hj, j-i)

			// compute area
			if t > a {
				a = t
			}
		}
	}
	return a
}

func maxAreaAllPairs2(height []int) int {
	maxarea := 0
	i := 0
	j := len(height) - 1
	for i < j {
		t := area(height[i], height[j], j-i)
		if t > maxarea {
			maxarea = t
		}

		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return maxarea
}

func maxInt(a, b int) int {
	if a < b {
		return b
	}
	return a
}

type data struct {
	lo int
	hi int
}

var dataMap map[data]int

var maxarea int

func maxR(height []int, lo int, hi int) int {
	var a int
	var b int
	var c int

	if hi <= lo {
		return 0
	}

	// d := data{hi: hi, lo: lo}
	// a, ok := dataMap[d]
	// if ok {
	// 	return a
	// }

	a = area(height[lo], height[hi], hi-lo)
	if a < maxarea {
		return maxarea
	}

	b = 0
	if height[lo] < height[hi] {
		b = maxR(height, lo+1, hi)
	}

	c = 0
	if height[hi] < height[lo] {
		c = maxR(height, lo, hi-1)
	}

	m := maxInt(a, maxInt(b, c))

	maxarea = maxInt(maxarea, m)

	return m
}

func maxAreaRecursive(height []int) int {
	dataMap = make(map[data]int)
	maxarea = 0
	b := maxR(height, 0, len(height)-1)
	return b
}
