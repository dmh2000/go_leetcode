package lc

func rangeBitWiseAnd(m int, n int) int {
	// if this is true
	if m*2 <= n {
		// all bits lower than the MSB of n will be 0
		// at some point in the count, zeroing out the
		// result
		return 0
	}
	// otherwise brute force it
	b := m
	for i := m + 1; i <= n; i++ {
		b = b & i
	}
	return b
}

func rangeBitWiseAnd2(m int, n int) int {
	// otherwise brute force it
	b := m
	for i := m + 1; i <= n; i++ {
		b = b & i
	}
	return b
}
