package lcmay

func pushBack1(q []int, v int) []int {
	q = append(q, v)
	return q
}

func popFront1(q []int) (int, []int) {
	var v int
	if len(q) == 0 {
		panic("popFront")
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

type digraph struct {
	v int     // number of vertices
	e int     // number of edges
	a [][]int // adjacency lists
}

func newDigraph(v int) *digraph {
	var g digraph

	g.v = v
	g.e = 0
	g.a = make([][]int, v)

	return &g
}

func (g *digraph) addEdge(v int, w int) {
	// create adjacency lists from edge list
	g.a[v] = append(g.a[v], w)
	g.e++
}

type directedCycle struct {
	marked  []bool // vertex was visited
	onstack []bool // vertex is in current path
	edgeTo  []int  // paths in dfs
	cycle   []int  // process queue
}

func newDirectedCycle(g *digraph) *directedCycle {
	var d directedCycle

	d.marked = make([]bool, g.v)
	d.onstack = make([]bool, g.v)
	d.edgeTo = make([]int, g.v)
	d.cycle = make([]int, 0)

	// look for cycle
	for v := 0; v < g.v; v++ {
		if !d.marked[v] && len(d.cycle) == 0 {
			d.dfs(g, v)
		}
	}

	return &d
}
func (d *directedCycle) dfs(g *digraph, v int) {
	d.onstack[v] = true
	d.marked[v] = true
	// for all vertices adjacent to v
	for _, w := range g.a[v] {
		// cycle queue should be 0 here or there's a cycle
		if len(d.cycle) > 0 {
			return
		}
		// if w has not been visited
		if !d.marked[w] {
			// add v as parent of w
			d.edgeTo[w] = v
			// continue search from w
			d.dfs(g, w)
		} else if d.onstack[w] {
			// enqueue all vertices that go from
			// v to w
			for x := v; x != w; x = d.edgeTo[x] {
				d.cycle = pushBack1(d.cycle, x)
			}
			// enqueue w
			d.cycle = pushBack1(d.cycle, w)
			// enqueue v
			d.cycle = pushBack1(d.cycle, v)
		}
	}
	d.onstack[v] = false
}

// =====================================================
func canFinish(numCourses int, prerequisites [][]int) bool {
	var g *digraph
	var d *directedCycle

	g = newDigraph(numCourses)
	for _, v := range prerequisites {
		g.addEdge(v[1], v[0])
	}

	d = newDirectedCycle(g)

	b := len(d.cycle) == 0
	return b
}
