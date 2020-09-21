package lcmay

func isPerfectSquare(num int) bool {

	if num == 0 {
		return false
	} else if num == 1 {
		return true
	}

	for i := 2; i <= num/2; i++ {
		sqr := i * i
		if sqr > num {
			return false
		} else if sqr == num {
			return true
		}

	}
	return false
}
