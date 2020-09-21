package lcmay

func iabs(n int) int {
	if n < 0 {
		n *= -1
	}
	return n
}

// slope is rise/run. points are on a line if slope is the same between all points
// integer division of rise/run might create a remainder so compute that also
// special case : if run is 0 the line is straight up and the division will fail
// so encode that with a value outside the range of the coordinates
func slope(p0 []int, p1 []int) (int, int) {
	var quo int
	var rem int
	// use positive values only, the results are the same but avoid sign issues
	run := iabs(p1[0] - p0[0])
	rise := iabs(p1[1] - p0[1])
	if run == 0 {
		// straigt up, slope is infinite
		quo = 2147483647
		rem = 2147483647
	} else {
		quo = rise / run
		rem = rise % run
	}

	// return both quotient and remainder
	return quo, rem
}

func checkStraightLine(coordinates [][]int) bool {

	q0, r0 := slope(coordinates[0], coordinates[1])
	for i := 1; i < len(coordinates); i++ {
		// both quotient and remainder must be the same for all points
		q, r := slope(coordinates[0], coordinates[i])
		if q != q0 || r != r0 {
			return false
		}
	}

	return true
}
