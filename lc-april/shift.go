package lc

func shiftLeft(s string, count int) string {

	for i := 0; i < count; i++ {
		c := s[0]
		s = s[1:] + string(c)
	}

	return s
}

func shiftRight(s string, count int) string {

	for i := 0; i < count; i++ {
		c := s[len(s)-1]
		s = string(c) + s[0:len(s)-1]
	}

	return s
}

func stringShift(s string, shift [][]int) string {

	for _, v := range shift {
		if v[0] == 0 {
			s = shiftLeft(s, v[1])
		} else {
			s = shiftRight(s, v[1])
		}
	}
	return s
}
