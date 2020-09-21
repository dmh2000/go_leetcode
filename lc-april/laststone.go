package lc

import (
	"container/heap"
)

// StoneHeap ...
type StoneHeap []int

func (h StoneHeap) Len() int           { return len(h) }
func (h StoneHeap) Less(i, j int) bool { return h[j] < h[i] }
func (h StoneHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push ...
func (h *StoneHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

// Pop ...
func (h *StoneHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func lastStoneWeight(stones []int) int {
	h := &StoneHeap{}
	heap.Init(h)
	for _, v := range stones {
		heap.Push(h, v)
	}

	var p int
	for {
		// if none left, return 0
		if h.Len() == 0 {
			p = 0
			break
		}
		// get largest value from top of heap
		p1 := heap.Pop(h).(int)

		// if last one, return its value
		if h.Len() == 0 {
			p = p1
			break
		}

		// smash them together and push the result
		p2 := heap.Pop(h).(int)
		if p1 > p2 {
			p = p1 - p2
			heap.Push(h, p)
		} else {
			p = p2 - p1
			heap.Push(h, p)
		}
	}

	return p
}
