package lcjune

import (
	"fmt"
	"testing"
)

func TestList1(t *testing.T) {

	var a []int
	var list *singleLinkedList
	var node *ListNode

	a = []int{1, 2, 3, 4}

	list = buildList(a)

	node = list.head
	for node != nil {
		fmt.Printf("%v,", node.Val)
		node = node.Next
	}

	node = list.front()
	if node.Val != 1 {
		t.Error(node.Val, " should be ", 1)
	}

	node = list.back()
	if node.Val != 4 {
		t.Error(node.Val, " should be ", 4)
	}

	i := 1
	for list.front() != nil {
		val, _ := list.popFront()
		if val != i {
			t.Error(val, " should be ", i)
		}
		i++
		if i > 4 {
			break
		}
	}
	if list.front() != nil {
		t.Error(list.front(), " shoudl be ", nil)
	}
}

func TestList2(t *testing.T) {

	var a []int
	var list *singleLinkedList
	var node *ListNode

	a = []int{}

	list = buildList(a)

	if list.size != 0 {
		t.Error(list.size, " should be ", 0)
	}

	for i := 4; i > 0; i-- {
		list.pushBack(i)
		if list.back().Val != i {
			t.Error(list.back().Val, " should be ", i)
		}
	}

	node = list.head
	for node != nil {
		fmt.Printf("%v,", node.Val)
		node = node.Next
	}

	node = list.front()
	if node.Val != 4 {
		t.Error(node.Val, " should be ", 4)
	}

	node = list.back()
	if node.Val != 1 {
		t.Error(node.Val, " should be ", 1)
	}

	i := 4
	for list.front() != nil {
		val, _ := list.popFront()
		if val != i {
			t.Error(val, " should be ", i)
		}
		i--
		if i == 0 {
			break
		}
	}
	if list.front() != nil {
		t.Error(list.front(), " shoudl be ", nil)
	}
}
