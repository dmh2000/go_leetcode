package lcmay

import (
	"sort"
)

func isAna1(a string, b string) bool {
	ma := make(map[byte]int)
	mb := make(map[byte]int)

	for i := 0; i < len(a); i++ {
		ma[a[i]]++
		mb[b[i]]++
	}

	for c := byte('a'); c <= byte('z'); c++ {
		if ma[c] != mb[c] {
			return false
		}
	}

	return true
}

func sortString(a string) string {
	b := []byte(a)
	sort.Slice(b, func(i int, j int) bool { return b[i] < b[j] })
	return string(b)
}

func insertUp(b []byte) []byte {
	for j := 0; j < len(b)-1; j++ {
		if b[j] > b[j+1] {
			b[j], b[j+1] = b[j+1], b[j]
		}
	}
	return b
}

func insertDown(b []byte) []byte {
	for j := len(b) - 1; j > 0; j-- {
		if b[j] < b[j-1] {
			b[j], b[j-1] = b[j-1], b[j]
		}
	}
	return b
}

func insertionByte(b []byte, r byte, c byte) []byte {
	// remove r from the string
	// and replace it with the new character
	var i int
	for i = 0; i < len(b); i++ {
		if b[i] == r {
			b[i] = c
			break
		}
	}

	if i == 0 {
		// insert up
		b = insertUp(b)
	} else if i == len(b)-1 {
		// insert down
		b = insertDown(b)
	} else if b[i] > b[i+1] {
		// insert up
		b = insertUp(b)
	} else {
		// insert down
		b = insertDown(b)
	}

	return b
}

func findAnagrams2(s string, p string) []int {
	var r []int
	var plen int
	var q string
	var sb []byte
	var pb []byte

	if len(s) == 0 {
		return []int{}
	}
	if len(p) == 0 {
		return []int{}
	}
	if len(p) > len(s) {
		return []int{}
	}

	sb = []byte(s)
	pb = []byte(p)

	r = make([]int, 0)

	// convert to byte slices to sort
	pb = []byte(p)
	plen = len(pb)
	// sort it
	sort.SliceStable(pb, func(i int, j int) bool { return pb[i] < pb[j] })
	// convert back to string
	p = string(pb)

	sb = []byte(s[0:plen])
	sort.SliceStable(sb, func(i int, j int) bool { return sb[i] < sb[j] })
	q = string(sb)
	for i := 0; i < len(s)-plen+1; i++ {
		// fmt.Println(p)
		// fmt.Println(q)
		// fmt.Println()
		if p == q {
			r = append(r, i)
		}
		next := i + plen
		if next >= len(s) {
			break
		}
		sb = insertionByte(sb, byte(s[i]), byte(s[next]))
		q = string(sb)
		// if !sort.SliceIsSorted(sb, func(i int, j int) bool { return sb[i] < sb[j] }) {
		// 	fmt.Println(q)
		// }
	}
	return r
}

func findAnagrams(s string, p string) []int {

	var mp []int
	var plen int
	var sp []int
	var r []int
	var b bool

	if len(s) == 0 {
		return []int{}
	}
	if len(p) == 0 {
		return []int{}
	}
	if len(p) > len(s) {
		return []int{}
	}

	r = make([]int, 0)
	plen = len(p)

	mp = make([]int, 26)
	for i := 0; i < plen; i++ {
		mp[p[i]-'a']++
	}

	sp = make([]int, 26)
	for j := 0; j < plen; j++ {
		sp[s[j]-'a']++
	}

	b = true
	for j := 0; j < 26; j++ {
		if mp[j] != sp[j] {
			b = false
		}
	}
	if b {
		r = append(r, 0)
	}

	for i := 1; i < len(s)-len(p)+1; i++ {
		sp[s[i-1]-'a']--
		// add one from sp[i+plen]
		sp[s[i+plen-1]-'a']++
		// check it
		b := true
		for j := 0; j < 26; j++ {
			if mp[j] != sp[j] {
				b = false
				break
			}
		}
		if b {
			r = append(r, i)
		}
	}

	return r
}
