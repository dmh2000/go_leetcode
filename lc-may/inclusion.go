package lcmay

import "fmt"

type freqCountLowerCase struct {
	count []int
	has   map[byte]bool
	total int
}

func newFreqCountLowerCase(s string) *freqCountLowerCase {
	var f freqCountLowerCase
	f.count = make([]int, 26)
	f.has = make(map[byte]bool)
	f.total = 0
	for i := 0; i < len(s); i++ {
		f.add(s[i])
		f.has[s[i]-'a'] = true
	}
	return &f
}

func (f *freqCountLowerCase) contains(c byte) bool {
	return f.has[c]
}

func (f *freqCountLowerCase) add(c byte) {
	f.count[c-'a']++
	f.total++
}

func (f *freqCountLowerCase) sub(c byte) {
	f.count[c-'a']--
	if f.count[c-'a'] < 0 {
		panic("count error less than 0")
	}
	f.total--
}

func (f *freqCountLowerCase) get(c byte) int {
	return f.count[c-'a']
}

func (f *freqCountLowerCase) reset(s string) {
	for i := 0; i < len(s); i++ {
		c := s[i]
		f.count[c-'a']++
		f.total++
	}

}

func (f *freqCountLowerCase) equals(g *freqCountLowerCase) bool {
	for i := 0; i < len(f.count); i++ {
		if f.count[i] != g.count[i] {
			return false
		}
	}
	return true
}

func (f *freqCountLowerCase) isZero() bool {
	var b bool

	b = true
	for i := 0; i < len(f.count); i++ {
		if f.count[i] != 0 {
			b = false
			break
		}
	}
	if b && f.total != 0 {
		panic("assert fail")
	}

	return b
}

func (f *freqCountLowerCase) print() {
	fmt.Println(f.total, f.count)
}

func checkInclusion(s1 string, s2 string) bool {

	if len(s1) > len(s2) {
		return false
	}

	if len(s1) == 1 && len(s2) == 1 && s1 == s2 {
		return true
	}

	// do frequency count for string 1
	var f *freqCountLowerCase
	var g *freqCountLowerCase

	f = newFreqCountLowerCase(s1)
	g = newFreqCountLowerCase(s2[0:len(s1)])
	if f.equals(g) {
		return true
	}
	k := 0
	for i := len(s1); i < len(s2); i++ {
		g.sub(s2[k])
		g.add(s2[i])
		if f.equals(g) {
			return true
		}
		k++
	}
	return false
}
