package lc

import (
	"container/heap"
)

// minHeap min heap of ints
type minHeap []int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push ...
func (h *minHeap) Push(x interface{}) {
	// add the
	*h = append(*h, x.(int))
}

// Pop ...
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Remove (extra, not in the heap interface)
func (h *minHeap) Remove(x interface{}) {
	var i int
	v := x.(int)
	for i = range *h {
		if (*h)[i] == v {
			break
		}
	}
	heap.Remove(h, i)
}

// ===============================================================
// stacklist
type intNode struct {
	value int
	next  *intNode
}

func (head *intNode) Push(v int) *intNode {
	node := &intNode{value: v, next: nil}
	node.next = head
	head = node
	return head
}

func (head *intNode) Pop() *intNode {
	if head == nil {
		panic(1)
	}
	head = head.next
	return head
}

func (head *intNode) Front() int {
	if head == nil {
		panic(1)
	}
	return head.value
}

// ===============================================================

// MinStack ...
type MinStack struct {
	stack *intNode
	heap  minHeap
}

/** initialize your data structure here. */

// NewMinStack ...
func NewMinStack() MinStack {
	mstack := MinStack{}
	heap.Init(&mstack.heap)
	return mstack
}

// Push ...
func (mstack *MinStack) Push(x int) {
	// slice the new value into the stack
	mstack.stack = mstack.stack.Push(x)

	// push it on the heap
	heap.Push(&mstack.heap, x)
}

// Pop ...
func (mstack *MinStack) Pop() {
	// get the top
	var top = mstack.stack.Front()

	// slice off the top
	mstack.stack = mstack.stack.Pop()

	// have to find the value in the heap then remove it
	// normal heap.Pop removes the min value, not the top
	mstack.heap.Remove(top)
}

// Top ...
func (mstack *MinStack) Top() int {
	return mstack.stack.Front()
}

// GetMin ...
func (mstack *MinStack) GetMin() int {
	return mstack.heap[0]
}

// Len ...
func (mstack *MinStack) Len() int {
	return mstack.heap.Len()
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := NewMinStack();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
