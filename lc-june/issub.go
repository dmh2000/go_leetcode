package lcjune

func isSubsequence(s string, t string) bool {
	var i int
	var j int

	if len(s) > len(t) {
		return false
	}

	if len(s) == 0 {
		return true
	}

	if len(t) == 0 {
		return false
	}

	j = 0
	for i = 0; i < len(s); i++ {
		// if s has k letters remaining
		// and j is shorter than that quit immediate
		k := len(s) - i
		m := len(t) - j
		if k > m {
			return false
		}

		// advance j until it matches s[i]
		for j < len(t) {
			// is there a match?
			if s[i] == t[j] {
				// matched last character of s, success
				if i == len(s)-1 {
					return true
				}
				// matched but more characters to go
				// advance j and break out of the loop
				j++
				break
			}
			// no match, just advance j
			j++
		}
	}
	return false
}
