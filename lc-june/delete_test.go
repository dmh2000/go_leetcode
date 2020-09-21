package lcjune

import "testing"

func TestDeleteNode1(t *testing.T) {
	var s *singleLinkedList

	s = buildList([]int{4, 5, 1, 9})
	printList(s.head)
	deleteNode(s.head)
	printList(s.head)

	s = buildList([]int{4, 5, 1, 9})
	printList(s.head)
	deleteNode(s.head.Next)
	printList(s.head)

	s = buildList([]int{4, 5, 1, 9})
	printList(s.head)
	deleteNode(s.head.Next.Next)
	printList(s.head)
}
