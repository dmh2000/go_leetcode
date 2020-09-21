package lcmay

import (
	"sort"
	"strings"
)

type freqPair struct {
	alph string
	freq int
}

type freqSlice []freqPair

func (fp freqSlice) Len() int {
	return len(fp)
}

func (fp freqSlice) Swap(i int, j int) {
	fp[i], fp[j] = fp[j], fp[i]
}

func (fp freqSlice) Less(i int, j int) bool {
	return fp[j].freq < fp[i].freq
}

func frequencySort1(s string) string {
	if len(s) <= 1 {
		return s
	}

	var f freqSlice
	f = make([]freqPair, 256)

	// for i := 0; i < len(f); i++ {
	// 	f[i].alph = ""
	// 	f[i].freq = 0
	// }

	for i := 0; i < len(s); i++ {
		f[s[i]].freq++
		f[s[i]].alph += string(s[i])
	}

	sort.Stable(f)

	r := ""
	for i := 0; i < len(f); i++ {
		if f[i].freq > 0 {
			r += f[i].alph
		}
	}

	return r
}

type sf []string

// type byteSlice []byte
func (a sf) Len() int           { return len(a) }
func (a sf) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sf) Less(i, j int) bool { return len(a[i]) > len(a[j]) }

func frequencySortS(s string) string {
	if len(s) <= 1 {
		return s
	}
	var a sf
	a = make([]string, 256)
	for i := 0; i < len(s); i++ {
		a[s[i]] += string(s[i])
	}

	sort.Sort(a)

	r := strings.Join(a, "")

	return r
}

type sb [][]byte

// type byteSlice []byte
func (a sb) Len() int           { return len(a) }
func (a sb) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sb) Less(i, j int) bool { return len(a[i]) > len(a[j]) }

// do all operations on the bytes of the string rather than
// on strings directly because concating strings is expensive
// give slices a large initial capacity so they don't reallocate frequently if at all
func frequencySortB(s string) string {
	if len(s) <= 1 {
		return s
	}
	var a sb
	var b []byte
	a = make([][]byte, 256, 10000)

	// build arrays of each byte
	for i := 0; i < len(s); i++ {
		a[s[i]] = append(a[s[i]], byte(s[i]))
	}

	// sort in order of decreasing length
	sort.Sort(a)

	// build the output byte array
	b = make([]byte, 0, 100000)
	for i := 0; i < len(a); i++ {
		b = append(b, a[i]...)
	}

	// convert it to string
	return string(b)
}
