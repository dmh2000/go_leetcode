package lc

// StackNode ...
type StackNode struct {
	next *StackNode
	data interface{}
}

// StackList ...
type StackList struct {
	head  *StackNode
	count int
}

// Stack ...
type Stack interface {
	PushFront(data interface{})
	PopFront(data interface{})
	Head() *StackNode
	Next(node *StackNode) *StackNode
	Empty() bool
}

// NewStackList ...
func NewStackList() *StackList {
	var list = StackList{
		head:  nil,
		count: 0,
	}
	return &list
}

// Head ...
func (list *StackList) Head() *StackNode {
	return list.head
}

// Next ...
func (list *StackList) Next(node *StackNode) *StackNode {
	return node.next
}

// Value ...
func (list *StackList) Value(node *StackNode) interface{} {
	return node.data
}

// PushFront ...
func (list *StackList) PushFront(data interface{}) {
	var node StackNode

	// data must be an int
	node.data = data

	// if empty, init the whole list
	if list.head == nil {
		list.head = &node
		list.count = 1
	} else {
		// not empty, push the node onto the head
		node.next = list.head
		list.head = &node
		list.count++
	}
}

// PopFront ...
func (list *StackList) PopFront() (interface{}, bool) {
	var next *StackNode
	if list.head == nil {
		return interface{}(0), false
	}

	// point at front node
	next = list.head

	// move head to next
	list.head = list.head.next

	return next.data, true
}

// Empty ...
func (list *StackList) Empty() bool {
	return list.head == nil
}
