package lcmay

import (
	"testing"
)

func TestOddEven1(t *testing.T) {
	var list *singleLinkedList
	var node *ListNode

	list = buildList([]int{1, 2, 3})
	node = oddEvenList(list.head)
	printList(node)

	list = buildList([]int{1, 2, 3, 4})
	node = oddEvenList(list.head)
	printList(node)

	list = buildList([]int{1, 2, 3, 4, 5})
	node = oddEvenList(list.head)
	printList(node)

	list = buildList([]int{})
	node = oddEvenList(list.head)
	printList(node)

	list = buildList([]int{1})
	node = oddEvenList(list.head)
	printList(node)

	list = buildList([]int{1, 2})
	node = oddEvenList(list.head)
	printList(node)

}
