package lc

// Graph ...
type Graph interface {
	// Graph : create a graph with no edges
	Graph(v int)

	// V : number of vertices
	V() int

	// E : number of edges
	E() int

	// AddEdge : add an edge from v to w
	AddEdge(v int, w int)

	//Adj : get a list of vertices adjacent to v
	Adj(v int) []int

	// ToString : return string representation
	ToString() string
}

// AdjacencyListGraph ...
type AdjacencyListGraph struct {
	v   int
	e   int
	adj [][]int
}

// NewAdjacencyListGraph : create and initialize
func NewAdjacencyListGraph(v int) *AdjacencyListGraph {
	ag := AdjacencyListGraph{}
	ag.Graph(v)
	return &ag
}

// Graph : create a graph with no edges
func (graph *AdjacencyListGraph) Graph(v int) {
	graph.v = v
	graph.e = 0
	graph.adj = make([][]int, v)
	// initialize the array of adjacency lists
	for i := range graph.adj {
		graph.adj[i] = make([]int, 0)
	}
}

// V : number of vertices
func (graph *AdjacencyListGraph) V() int {
	return graph.v
}

// E : number of edges
func (graph *AdjacencyListGraph) E() int {
	return graph.e
}

// AddEdge : add an edge from v to w
func (graph *AdjacencyListGraph) AddEdge(v int, w int) {
	graph.adj[v] = append(graph.adj[v], w)
	//graph.adj[w] = append(graph.adj[w], v)
}

//Adj : get a list of vertices adjacent to v
func (graph *AdjacencyListGraph) Adj(v int) []int {
	return graph.adj[v]
}

// ToString : return string representation
func (graph *AdjacencyListGraph) ToString() string {
	return "AdjacencyListGraph"
}

// CC ...
type CC struct {
	marked []bool
	id     []int
	size   []int
	count  int
}

func ccDfs(cc *CC, g *AdjacencyListGraph, v int) {
	cc.marked[v] = true
	cc.id[v] = cc.count
	cc.size[cc.count]++
	for _, w := range g.Adj(v) {
		if !cc.marked[w] {
			ccDfs(cc, g, w)
		}
	}
}

func ccCount(cc *CC) int {
	return cc.count
}

// NewCC ...
func NewCC(g *AdjacencyListGraph) *CC {
	cc := CC{}
	cc.marked = make([]bool, g.V())
	cc.id = make([]int, g.V())
	cc.size = make([]int, g.V())
	for v := 0; v < g.V(); v++ {
		if !cc.marked[v] {
			ccDfs(&cc, g, v)
			cc.count++
		}
	}
	return &cc
}

func numIslands(grid [][]byte) int {
	var g *AdjacencyListGraph
	var v int
	vertices := len(grid) * len(grid[0])
	g = NewAdjacencyListGraph(vertices)

	id := make([][]int, 0)
	v = 0
	for i, rv := range grid {
		id = append(id, []int{})
		for range rv {
			id[i] = append(id[i], v)
			v++
		}
	}

	v = 0
	z := 0
	for r, rv := range grid {
		for c, cw := range rv {
			v++

			// skip empty cells
			if cw == 0 {
				z++
				continue
			}
			// left
			if c > 0 && rv[c-1] == 1 {
				g.AddEdge(id[r][c], id[r][c-1])
			}
			// right
			if c < len(rv)-1 && rv[c+1] == 1 {
				g.AddEdge(id[r][c], id[r][c+1])
			}
			// up
			if r > 0 && grid[r-1][c] == 1 {
				g.AddEdge(id[r][c], id[r-1][c])
			}
			// down
			if r < len(grid)-1 && grid[r+1][c] == 1 {
				g.AddEdge(id[r][c], id[r+1][c])
			}
		}
	}
	// for v := 0; v < vertices; v++ {
	// 	fmt.Printf("%v : ", v)
	// 	for _, w := range g.Adj(v) {
	// 		fmt.Printf("%v ", w)
	// 	}
	// 	fmt.Printf("\n")
	// }

	var cc *CC
	cc = NewCC(g)
	return ccCount(cc) - z
}
