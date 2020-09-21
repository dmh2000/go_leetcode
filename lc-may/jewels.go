package lcmay

func numJewelsInStones1(J string, S string) int {
	jewels := make(map[rune]int)

	// create a map with only the jewels in it
	for _, v := range J {
		jewels[v] = 0
	}

	// update counts for jewels only
	for _, v := range S {
		_, ok := jewels[v]
		if ok {
			jewels[v]++
		}
	}

	// iterate through the map for jewel keys and sum them up
	sum := 0
	for _, v := range jewels {
		sum += v
	}

	return sum
}

// instead of a map, use a array indexed by the individual runes
func numJewelsInStones2(J string, S string) int {
	var a []int

	a = make([]int, 255)

	// mark the array with -1 to indicate not present
	for i := range a {
		a[i] = -1
	}

	// set jewels to 1 to indicate present
	for _, r := range J {
		a[int(r)] = 1
	}

	// update the array with the stones
	sum := 0
	for _, r := range S {
		if a[int(r)] == 1 {
			sum++
		}
	}

	return sum
}

func numJewelsInStones3(J string, S string) int {
	jewels := make(map[rune]int)

	// create a map with only the jewels in it
	for _, v := range J {
		jewels[v] = 0
	}

	// update counts for jewels only
	sum := 0
	for _, v := range S {
		_, ok := jewels[v]
		if ok {
			sum++
		}
	}

	return sum
}
