package lcmay

import "container/heap"

type kval struct {
	d int
	p []int
}
type kheap []kval

func (h kheap) Len() int           { return len(h) }
func (h kheap) Less(i, j int) bool { return h[i].d < h[j].d }
func (h kheap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *kheap) Push(x interface{}) {
	*h = append(*h, x.(kval))
}

func (h *kheap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func kdistance(p []int) int {
	return p[0]*p[0] + p[1]*p[1]
}

func kClosest(points [][]int, K int) [][]int {
	var h *kheap
	var kv kval
	var a [][]int
	h = &kheap{}
	heap.Init(h)

	// compute the distances
	for _, v := range points {
		kv.p = v
		kv.d = kdistance(v)
		heap.Push(h, kv)
	}

	// get the k closest
	a = make([][]int, 0)
	for i := 0; i < K; i++ {
		kv = heap.Pop(h).(kval)
		a = append(a, kv.p)
	}

	return a
}
