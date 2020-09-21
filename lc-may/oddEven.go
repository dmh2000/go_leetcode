package lcmay

func oddEvenList(head *ListNode) *ListNode {
	var oddhead *ListNode
	var evenhead *ListNode
	var odd *ListNode
	var even *ListNode
	var tail *ListNode
	var n *ListNode
	var b int

	if head == nil {
		return nil
	}

	// split the lists
	b = 0
	n = head

	oddhead = nil
	evenhead = nil

	for head != nil {
		// detach a first node
		n = head
		head = head.Next
		n.Next = nil

		if b == 0 {
			if oddhead == nil {
				oddhead = n
				odd = n
				tail = odd
			} else {
				odd.Next = n
				odd = odd.Next
				tail = odd
			}
			b = 1
		} else {
			if evenhead == nil {
				evenhead = n
				even = n
			} else {
				even.Next = n
				even = even.Next
			}
			b = 0
		}
	}

	if odd != nil {
		tail.Next = evenhead
	}
	return oddhead
}
