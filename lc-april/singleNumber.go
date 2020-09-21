package lc

import (
	"container/heap"
)

type xh struct {
	number int
	count  int
}

type xheap []xh

func (h xheap) Len() int {
	return len(h)
}

func (h xheap) Less(i, j int) bool {
	return h[i].count < h[j].count
}

func (h xheap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *xheap) Push(x interface{}) {
	*h = append(*h, x.(xh))
}

func (h *xheap) Pop() interface{} {
	old := *h
	n := len(old)
	a := old[n-1]
	*h = old[0 : n-1]
	return a
}

func singleNumberHeap(nums []int) int {
	m := make(map[int]int)

	// build the hash table
	for _, v := range nums {
		m[v] = m[v] + 1
	}

	// now add to heap
	h := &xheap{}
	for k, v := range m {
		a := xh{number: k, count: v}
		heap.Init(h)
		heap.Push(h, a)
	}

	b := heap.Pop(h)
	number := b.(xh).number

	return number
}

func singleNumberMap(nums []int) int {
	m := make(map[int]int)

	// build the hash table
	for _, v := range nums {
		m[v] = m[v] + 1
	}

	// scan throug map to find count of 1
	var n int
	for k, v := range m {
		if v == 1 {
			n = k
			break
		}
	}
	return n
}

func singleNumberMap2(nums []int) int {
	m := make(map[int]int)

	// build the hash table
	for _, v := range nums {
		_, ok := m[v]
		if ok {
			delete(m, v)
		} else {
			m[v] = v
		}
	}

	for _, v := range m {
		return v
	}
	return 0
}

func singleNumberMapDiv(nums []int) int {

	s := 0
	for _, v := range nums {
		s ^= v
	}

	return s
}

// C++
// class Solution {
// 	public:
// 		int singleNumber(vector<int>& nums) {
// 			int s = 0;
// 			for (int i=0;i<nums.size();i++) {
// 				s ^= nums[i];
// 			}
// 			return s;
// 		}
// 	};
