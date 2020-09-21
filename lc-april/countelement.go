package lc

import (
	"sort"
)

func countElements(arr []int) int {
	sort.IntSlice.Sort(arr)

	count := 0
	for i := 0; i < len(arr); i++ {
		// process the first digit
		d := arr[i]
		x := d + 1
		k := sort.Search(len(arr), func(j int) bool { return arr[j] >= x })
		if k < len(arr) && arr[k] == d+1 {
			// x is  found
			count++
		}
	}
	return count
}
