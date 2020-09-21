package lcjune

type step struct {
	pos  int
	k    int
	cost int
}

type loc struct {
	next int
	cost int
}

func findCheapestPrice(
	n int,
	flights [][]int,
	src int,
	dst int,
	k int) int {

	var q []step
	var price map[int]int
	var location map[int][]loc

	price = make(map[int]int)
	for i := 0; i < n; i++ {
		price[i] = 2147483648
	}

	location = make(map[int][]loc)
	for _, v := range flights {
		key := v[0]
		val := loc{v[1], v[2]}
		location[key] = append(location[key], val)
	}

	// process queue
	q = []step{}

	// start with source
	q = append(q, step{src, -1, 0})

	for len(q) > 0 {
		// pop off the queue
		s := q[0]
		if len(q) > 0 {
			q = q[1:]
		}

		// too many stops
		if s.pos == dst || s.k == k {
			continue
		}

		for _, loc := range location[s.pos] {
			next := loc.next
			cost := loc.cost + s.cost
			if cost <= price[next] {
				price[next] = cost
				q = append(q, step{next, s.k + 1, cost})
			}
		}

	}
	p := price[dst]
	if p == 2147483648 {
		p = -1
	}
	return p
}
