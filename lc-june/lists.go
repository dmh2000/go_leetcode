package lcjune

import "fmt"

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

type singleLinkedList struct {
	head *ListNode
	tail *ListNode
	size int
}

func (list *singleLinkedList) pushBack(val int) {
	var node ListNode

	// init the node
	node = ListNode{val, nil}

	// if its empty
	if list.head == nil {
		list.head = &node
		list.tail = &node
		list.size = 1
		return
	}

	// add it on the tail
	list.tail.Next = &node
	list.tail = list.tail.Next
	list.size++
}

func (list *singleLinkedList) pushFront(val int) {
	var node ListNode

	// init the node
	node = ListNode{val, nil}

	// if its empty
	if list.head == nil {
		list.head = &node
		list.tail = &node
		list.size = 1
		return
	}

	// add it at the head
	node.Next = list.head
	list.head = &node

	list.size++
}

func (list *singleLinkedList) popFront() (int, bool) {
	var val int

	if list.head == nil {
		return 0, false
	}

	// init the node
	val = list.head.Val

	// remove the head
	list.head = list.head.Next

	// update tail if list is empty
	if list.head == nil {
		list.tail = nil
	}

	// one less item
	list.size--

	return val, true
}

func (list *singleLinkedList) front() *ListNode {
	return list.head
}

func (list *singleLinkedList) back() *ListNode {
	return list.tail
}

// build singly linked list
func buildList(a []int) *singleLinkedList {
	var list singleLinkedList
	if len(a) == 0 {
		list.head = nil
		list.tail = nil
		list.size = 0
		return &list
	}

	list.head = &ListNode{a[0], nil}
	list.tail = list.head
	list.size = 1

	for i := 1; i < len(a); i++ {
		list.pushBack(a[i])
	}

	return &list
}

func printList(n *ListNode) {
	for n != nil {
		fmt.Printf("%v,", n.Val)
		n = n.Next
	}
	fmt.Println()
}
