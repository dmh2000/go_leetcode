package lc

import "fmt"

// IntListNode ...
type IntListNode struct {
	next  *IntListNode
	prev  *IntListNode
	value int
}

// IntList ...
type IntList struct {
	head  *IntListNode
	tail  *IntListNode
	count int
}

// NewIntList ...
func NewIntList() *IntList {
	var list = IntList{
		head:  nil,
		tail:  nil,
		count: 0,
	}
	return &list
}

// Front ...
func (list *IntList) Front() *IntListNode {
	return list.head
}

// Next ...
func (list *IntList) Next(node *IntListNode) *IntListNode {
	return node.next
}

// Prev ...
func (list *IntList) Prev(node *IntListNode) *IntListNode {
	return node.prev
}

// Tail ...
func (list *IntList) Tail() *IntListNode {
	return list.tail
}

// Value ...
func (list *IntList) Value(node *IntListNode) int {
	return node.value
}

// PushFront ...
func (list *IntList) PushFront(value int) *IntListNode {
	var node IntListNode
	node.next = nil
	node.prev = nil
	node.value = value

	if list.head == nil {
		list.head = &node
		list.tail = &node
		list.count = 1
	} else {
		// not empty,
		node.next = list.head
		node.next.prev = &node
		node.prev = nil
		list.head = &node
		list.count++
	}
	return &node
}

// PopFront ...
func (list *IntList) PopFront() (int, bool) {
	var value int
	if list.head == nil {
		return 0, false
	}

	// get front node value
	value = list.head.value

	if list.head == list.tail {
		// at head
		list.head = nil
		list.tail = nil
	} else {
		// move head to next
		list.head = list.head.next
		list.head.prev = nil
	}
	// decrement length
	list.count--

	return value, true
}

// PushBack ...
func (list *IntList) PushBack(value int) *IntListNode {
	var node IntListNode
	node.next = nil
	node.prev = nil
	node.value = value

	if list.head == nil {
		list.head = &node
		list.tail = &node
		list.count = 1
	} else {
		// not empty,
		list.tail.next = &node
		node.prev = list.tail
		list.tail = &node
		list.count++
	}
	return &node
}

// PopBack ...
func (list *IntList) PopBack() (int, bool) {
	var value int
	if list.head == nil {
		return 0, false
	}

	if list.tail == list.head {
		// at head
		return list.PopFront()
	}

	// not at head
	value = list.tail.value
	prev := list.tail.prev
	prev.next = nil
	list.tail = prev

	// decrement length
	list.count--

	return value, true
}

// Remove ...
func (list *IntList) Remove(node *IntListNode) {
	if node == nil || list.head == nil {
		panic("node is nil")
	}
	if node == list.head {
		// at head, advance head to next node
		list.head = node.next
		list.head.prev = nil
	} else if node == list.tail {
		// at tail
		list.tail = node.prev
		list.tail.next = nil
	} else {
		// in between
		prev := node.prev
		next := node.next
		prev.next = next
		next.prev = prev
	}
	// decrement count
	list.count--
}

// Len ...
func (list *IntList) Len() int {
	return list.count
}

// Empty ...
func (list *IntList) Empty() bool {
	return list.head == nil
}

// PrintFwd ...
func (list *IntList) PrintFwd() {
	node := list.head
	for node != nil {
		fmt.Printf("%v,", node)
		node = list.Next(node)
	}
	fmt.Println()
}

// PrintBak ...
func (list *IntList) PrintBak() {
	node := list.tail
	for node != nil {
		fmt.Printf("%v,", node)
		node = list.Prev(node)
	}
	fmt.Println()
}
