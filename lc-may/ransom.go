package lcmay

func canConstruct(ransomNote string, magazine string) bool {
	r := make([]int, 26)
	m := make([]int, 26)

	// frequency count ransom
	for i := 0; i < len(ransomNote); i++ {
		c := ransomNote[i] - 'a'
		r[c]++
	}

	// frequency count magazine
	for i := 0; i < len(magazine); i++ {
		c := magazine[i] - 'a'
		m[c]++
	}

	// make sure magazine has enough characters
	b := true
	for i := 0; i < 26; i++ {
		if r[i] > m[i] {
			// not enough characters
			b = false
			break
		}
	}

	return b
}
