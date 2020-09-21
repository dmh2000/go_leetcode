package lc

func isPalindrome(s string) bool {
	var pal bool

	pal = false
	i := 0
	j := len(s) - 1
	for s[i] == s[j] {
		i++
		j--
		if i >= j {
			pal = true
			break
		}
	}
	return pal
}

func isPalindromeSlice(s []byte) bool {
	var pal bool

	pal = false
	i := 0
	j := len(s) - 1
	for s[i] == s[j] {
		i++
		j--
		if i >= j {
			pal = true
			break
		}
	}
	return pal
}

func pal(s string) string {
	if len(s) == 1 {
		return s
	}

	if isPalindrome(s) {
		return s
	}

	big := ""
	for i := 0; i < len(s); i++ {
		if len(big) > (len(s) - i) {
			break
		}
		for j := len(s); j > i; j-- {
			if (j - i) < len(big) {
				break
			}
			a := s[i:j]
			if isPalindrome(a) {
				if len(a) > len(big) {
					big = a
					break
				}
			}
		}
	}

	return big
}

func palSlice(s []byte) string {
	if len(s) == 1 {
		return string(s)
	}

	if isPalindromeSlice(s) {
		return string(s)
	}

	big := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if len(big) > (len(s) - i) {
			break
		}
		for j := len(s); j > i; j-- {
			if (j - i) < len(big) {
				break
			}
			a := s[i:j]
			if isPalindromeSlice(a) {
				if len(a) > len(big) {
					big = a
					break
				}
			}
		}
	}

	return string(big)
}

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	a := pal(s)
	// a := palSlice([]byte(s))

	return a
}
