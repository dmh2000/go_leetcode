package lc

import "sort"

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Len() int {
	return len(s)
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func sortString(str string) string {
	// convert to runes
	s := []rune(str)
	sort.Sort(sortRunes(s))
	return string(s)
}

func groupAnagrams(strs []string) [][]string {
	var r [][]string
	r = make([][]string, 0)

	// sort the names in the string
	m := make(map[string][]string)
	for i := 0; i < len(strs); i++ {
		s := sortString(strs[i])
		m[s] = append(m[s], strs[i])
	}

	for _, v := range m {
		r = append(r, v)
	}

	return r
}
