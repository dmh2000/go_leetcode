package lcjune

import (
	"sort"
)

type rslice [][]int

func (r rslice) Len() int { return len(r) }

func (r rslice) Less(i int, j int) bool {
	a := r[i]
	b := r[j]
	if b[0] < a[0] {
		return true
	} else if a[0] == b[0] {
		if a[1] < b[1] {
			return true
		}
		return false
	}
	return false
}

func (r rslice) Swap(i int, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r rslice) InsertK(i int, b []int) {
	j := i
	k := b[1]
	r[i] = b
	for j > k {
		r.Swap(j, j-1)
		j--
	}
}

func reconstructQueue(people [][]int) [][]int {
	var a [][]int

	// quick check just in case
	if len(people) <= 1 {
		return people
	}

	// create destination slice
	a = make([][]int, len(people))

	// sort descending by [0], ascending by [1]
	sort.Sort(rslice(people))

	// insert each into slice from largest to smallest
	for i := 0; i < len(people); i++ {
		// place next element into [i]
		rslice(a).InsertK(i, people[i])
	}
	return a
}
