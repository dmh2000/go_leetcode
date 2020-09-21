package lcjune

import (
	"container/heap"
)

// =======================================
// minpq
// =======================================
type pnode struct {
	cost  int
	count int
}
type minpq []int

func (h minpq) Len() int           { return len(h) }
func (h minpq) Less(i, j int) bool { return h[i] < h[j] }
func (h minpq) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minpq) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *minpq) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// =======================================
// dijkstra
// =======================================

func dijkstra(g *edgeWeightedDigraph, src int, tgt int) int {
	var h *minpq
	var dist []int
	var step []int
	var prev []int

	dist = make([]int, g.V())
	prev = make([]int, g.V())
	step = make([]int, g.V())

	h = &minpq{}
	heap.Init(h)
	for v := 0; v < len(dist); v++ {
		dist[v] = 2147483648
		prev[v] = -1
		heap.Push(h, v)
	}
	dist[src] = 0

	// init the heap
	for h.Len() > 0 {
		// get min distance from source
		v := heap.Pop(h).(int)

		if v == tgt {
			break
		}

		// for all neighbors of v
		a := g.adj(v)
		for _, e := range a {
			w := e.to()
			alt := dist[v] + e.weight()
			if alt < dist[w] {
				step[w]++
				dist[w] = alt
				prev[w] = v
			}
		}
	}
	return dist[tgt]
}

func execDijkstra(
	n int,
	edges [][]int,
	src int,
	dst int,
) int {

	var g *edgeWeightedDigraph

	// create the graph
	g = newEdgeWeightedDigraph(n)

	// add all the edges
	for _, f := range edges {
		var e directedEdge
		e.v = f[0]
		e.w = f[1]
		e.cost = f[2]
		g.addEdge(e)
	}

	cost := dijkstra(g, src, dst)

	return cost
}

type edgeWeightedDigraph struct {
	v int              // number of vertices
	e int              // number of edges
	a [][]directedEdge // adjacency lists
}

type directedEdge struct {
	v    int // from
	w    int // to
	cost int
}

// ==========================
// directedEdge
// ==========================
func newDirectedEdge(v int, w int, cost int) directedEdge {
	var e directedEdge
	e.v = v
	e.w = w
	e.cost = cost
	return e
}

func (e directedEdge) from() int {
	return e.v
}

func (e directedEdge) to() int {
	return e.w
}

func (e directedEdge) weight() int {
	return e.cost
}

// ==========================
// edgeWeightedDigraph
// ==========================

func newEdgeWeightedDigraph(v int) *edgeWeightedDigraph {
	var g edgeWeightedDigraph

	g.v = v
	g.e = 0
	g.a = make([][]directedEdge, v)
	for i := 0; i < len(g.a); i++ {
		g.a[i] = make([]directedEdge, 0)
	}

	return &g
}

func (g *edgeWeightedDigraph) addEdge(e directedEdge) {
	// create adjacency lists from edge list
	g.a[e.from()] = append(g.a[e.from()], e)
	g.e++
}

func (g *edgeWeightedDigraph) V() int {
	return g.v
}

func (g *edgeWeightedDigraph) E() int {
	return g.e
}

func (g *edgeWeightedDigraph) adj(v int) []directedEdge {
	return g.a[v]
}
