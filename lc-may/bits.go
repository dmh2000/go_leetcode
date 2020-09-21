package lcmay

func bitsC(n int) int {
	count := 0
	for n > 0 {
		if n&0x01 == 0x01 {
			count++
		}
		n = n >> 1
	}
	return count
}

func bitsK(n int) int {
	var c int
	for c = 0; n > 0; n = n & (n - 1) {
		c++
	}
	return c
}

func bitsD(n int) int {
	n = n - ((n >> 1) & 0x55555555)
	n = (n & 0x33333333) + ((n >> 2) & 0x33333333)
	n = (n + (n >> 4)) & 0x0F0F0F0F
	n = n + (n >> 8)
	n = n + (n >> 16)
	return n & 0x0000003F
}

func countBits(num int) []int {
	var a []int

	a = make([]int, 0)
	for n := 0; n <= num; n++ {
		b := bitsK(n)
		a = append(a, b)
	}
	return a
}

func countBitsC(num int) []int {
	var a []int

	a = make([]int, 0)
	for n := 0; n <= num; n++ {
		b := bitsC(n)
		a = append(a, b)
	}
	return a
}

func countBitsD(num int) []int {
	var a []int

	a = make([]int, 0)
	for n := 0; n <= num; n++ {
		b := bitsD(n)
		a = append(a, b)
	}
	return a
}
