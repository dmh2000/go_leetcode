package lcmay

type uniquePair1 struct {
	index int
	count int
}

func firstUniqChar1(s string) int {
	m := make(map[rune]uniquePair1)

	for i, v := range s {
		_, ok := m[v]
		if ok {
			// already counted
			// don't set index
			u := m[v]
			u.count++
			m[v] = u
		} else {
			// first encounter, set the index
			u := m[v]
			u.count++
			u.index = i
			m[v] = u
		}
	}

	i := 2147483647
	c := rune(0)
	for k, u := range m {
		if u.count == 1 {
			if u.index < i {
				i = u.index
				c = k
			}
		}
	}
	// not found
	if c == 0 {
		i = -1
	}

	return i
}

type uniquePair struct {
	index int
	count int
}

func firstUniqChar(s string) int {
	a := make([]uniquePair, 26)

	for i := 0; i < 26; i++ {
		a[i].index = -1
	}

	// frequency count
	for i := 0; i < len(s); i++ {
		v := s[i] - 'a'
		a[v].count++
		if a[v].index < 0 {
			a[v].index = i
		}
	}

	idx := 2147483647
	for i := 0; i < 26; i++ {
		if a[i].count == 1 {
			if a[i].index < idx {
				idx = a[i].index
			}
		}
	}
	if idx == 2147483647 {
		idx = -1
	}

	return idx
}
