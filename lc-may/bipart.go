package lcmay

const noc = 0
const red = 1
const grn = 2

func pushBack(q []int, v int) []int {
	q = append(q, v)
	return q
}

func pushFront(q []int, v int) []int {
	q = append([]int{v}, q...)
	return q
}

func popBack(q []int) (int, []int) {
	var v int
	if len(q) == 0 {
		return -1, q
	}

	// get back
	v = q[len(q)-1]

	// remove back
	if len(q) == 1 {
		// clear the queue
		q = q[:0]
	} else {
		// remove element
		q = q[0 : len(q)-1]
	}

	return v, q
}

func popFront(q []int) (int, []int) {
	var v int
	if len(q) == 0 {
		return -1, q
	}

	// get front
	v = q[0]

	// remove front
	if len(q) == 1 {
		// clear the queue
		q = make([]int, 0)
	} else {
		// remove element
		q = q[1:]
	}

	return v, q
}
func connected(g [][]int) []int {
	var visited []bool
	var q []int
	var comps []int
	var u int
	var v int

	visited = make([]bool, len(g))
	comps = make([]int, 0, len(g))
	q = make([]int, 0, len(g))
	for u = 1; u < len(g); u++ {
		if !visited[u] {
			// start of component
			comps = append(comps, u)
			q = pushBack(q, u)
			visited[u] = true
			for len(q) != 0 {
				v, q = popBack(q)
				for w := 1; w < len(g); w++ {
					if g[v][w] == 1 {
						if !visited[w] {
							visited[w] = true
							q = pushBack(q, w)
						}
					}
				}
			}
		}
	}
	return comps
}

func possibleBipartition1(N int, dislikes [][]int) bool {
	var b bool
	var c []int
	var g [][]int
	var q []int
	var comps []int

	g = make([][]int, N+1)
	for i := 1; i < N+1; i++ {
		g[i] = make([]int, N+1)
	}

	for i := 0; i < len(dislikes); i++ {
		d := dislikes[i]
		u := d[0]
		v := d[1]
		g[u][v] = 1
		g[v][u] = 1
	}

	// get connected components
	comps = connected(g)

	c = make([]int, N+1)
	q = make([]int, 0, N+1)
	b = true
loop:
	for _, comp := range comps {

		q = pushBack(q, comp)
		c[comp] = red
		for len(q) != 0 {
			var u int

			// pop next vertex
			u, q = popFront(q)

			// for each edge out of u
			for v := 1; v < len(g); v++ {
				// self loop?
				if g[u][u] != noc {
					b = false
					break loop
				}

				// is there an edge from u to v
				if g[u][v] != 0 {
					if c[v] == noc {
						if c[u] == red {
							c[v] = grn
						} else {
							c[v] = red
						}
						q = pushBack(q, v)
					} else {
						// all ready visted
						if c[u] == c[v] {
							b = false
							break loop
						}
					}
				}
			}
		}
	}

	return b
}

func possibleBipartition(N int, dislikes [][]int) bool {
	var a [][]int
	var stack []int
	var seen map[int]int
	var b bool
	var exists bool

	a = make([][]int, N+1)
	for i := 0; i < len(a); i++ {
		a[i] = make([]int, 0)
	}

	for _, d := range dislikes {
		i := d[0]
		j := d[1]
		a[i] = append(a[i], j)
		a[j] = append(a[j], i)
	}

	seen = make(map[int]int)

	b = true
loop:
	for u := 1; u < N+1; u++ {
		// has i already been processed?
		_, exists = seen[u]
		if !exists {
			// no, put it in room 0 and put it on stack
			seen[u] = 0
			stack = pushFront(stack, u)
			// process the stack until empty
			for len(stack) > 0 {
				var v int
				v, stack = popFront(stack)
				// for all haters in list for v
				for _, w := range a[v] {
					_, exists = seen[w]
					if !exists {
						seen[w] = (seen[v] + 1) % 2
						stack = pushFront(stack, w)
					} else if seen[v] == seen[w] {
						b = false
						break loop
					}
				}
			}
		}
	}
	return b
}
