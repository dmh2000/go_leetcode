package lcmay

func ilog2(n int) int {
	if n < 2 {
		return 1
	}
	m := 2
	log := 0
	for m <= n {
		m = m * 2
		log++
	}
	return log
}

func findComplement(num int) int {

	if num == 0 {
		return 1
	}
	if num == 1 {
		return 0
	}

	// compute number of bits in result
	// f := math.Log2(float64(num))
	// nbits := uint32(f)
	nbits := ilog2(num)

	// create a mask for that number of bits
	mask := uint32(0xffffffff << nbits)
	mask = ^mask
	// fmt.Printf("%032b %v\n", mask, nbits)
	// invert num and mask its upper bits
	inv := ^num & int(mask)

	return inv
}
