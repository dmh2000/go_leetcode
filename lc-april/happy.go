package lc

// parse a number into its digits
func digits(n int) []int {
	d := make([]int, 0)
	exp := 10
	for n > 0 {
		v := n % exp
		d = append(d, v)
		n /= exp
	}
	return d
}

func square(d []int) int {
	sum := 0
	for _, v := range d {
		sum += v * v
	}
	return sum
}

func isHappy(n int) bool {
	var history map[int]bool
	var m int
	var happy bool
	var d []int

	history = make(map[int]bool, 0)
	happy = false
	m = n
	for true {
		d = digits(m)
		m = square(d)
		if m == 1 {
			happy = true
			break
		} else if history[m] {
			happy = false
			break
		}
		history[m] = true
	}
	return happy
}
