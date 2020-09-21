package lcjune

import (
	"sort"
)

type diffSlice [][]int

func (d diffSlice) Len() int {
	return len(d)
}

func (d diffSlice) Less(i int, j int) bool {
	return d[i][0] < d[j][0]
}

func (d diffSlice) Swap(i int, j int) {
	d[i], d[j] = d[j], d[i]
}

func twoCitySchedCost(costs [][]int) int {
	var cost int
	var diff [][]int

	cost = 0
	diff = make([][]int, len(costs))
	// for all the elements
	for i := 0; i < len(costs); i++ {
		var j int
		j = costs[i][0] - costs[i][1]
		// record the difference and the costs
		diff[i] = []int{j, costs[i][0], costs[i][1]}
	}

	// sort by difference in cost
	// this moves all the A's to the first half
	// and all the B's to the second half
	// that is because the differences (A-B cost)
	//  in the first half will be negative
	// and in the second half they will be positive
	sort.Sort(diffSlice(diff))

	// add up the A's and the B's
	cost = 0
	for i := 0; i < len(costs); i++ {
		if i < len(costs)/2 {
			// A's cost
			cost += diff[i][1]
		} else {
			// B's cost
			cost += diff[i][2]
		}
	}

	return cost
}
