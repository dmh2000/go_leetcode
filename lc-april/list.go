package lc

// MyListNode ...
type MyListNode struct {
	next *MyListNode
	data interface{}
}

// List ...
type List struct {
	head  *MyListNode
	tail  *MyListNode
	count int
}

// LinkedList ...
type LinkedList interface {
	PushFront(data interface{})
	PopFront(data interface{})
	// PushBack(data interface)
	// PopBack(data interface)
	Head() *MyListNode
	Next(node *MyListNode) *MyListNode
}

// MyIntList ...
type MyIntList List

// NewMyIntList ...
func NewMyIntList() *MyIntList {
	var list = MyIntList{
		head:  nil,
		tail:  nil,
		count: 0,
	}
	return &list
}

// Head ...
func (list *MyIntList) Head() *MyListNode {
	return list.head
}

// Next ...
func (list *MyIntList) Next(node *MyListNode) *MyListNode {
	return node.next
}

// Tail ...
func (list *MyIntList) Tail() *MyListNode {
	return list.tail
}

// Value ...
func (list *MyIntList) Value(node *MyListNode) int {
	return node.data.(int)
}

// PushFront ...
func (list *MyIntList) PushFront(data int) {
	var node MyListNode

	// data must be an int
	node.data = data

	// if empty, init the whole list
	if list.head == nil {
		list.head = &node
		list.tail = &node
		list.count = 1
	} else {
		// not empty, push the node onto the head
		node.next = list.head
		list.head = &node
		list.count++
	}
}

// PopFront ...
func (list *MyIntList) PopFront() (int, bool) {
	var next *MyListNode
	if list.head == nil {
		return 0, false
	}

	// point at front node
	next = list.head

	// move head to next
	list.head = list.head.next

	// if at end of list, set tail to nil also
	if list.head == nil {
		list.tail = nil
	}
	return next.data.(int), true
}
