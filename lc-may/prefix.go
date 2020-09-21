package lcmay

func longestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	// first get shortest length
	n := 2147483647
	for _, v := range strs {
		if len(v) < n {
			n = len(v)
		}
	}

	prefix := ""
	for i := 0; i < n; i++ {
		c := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if strs[j][i] != c {
				// mismatch, current prefix is the shortest
				return prefix
			}
		}
		prefix += string(c)
	}
	return prefix
}
