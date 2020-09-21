package lc

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	n := head
	m := head
	// for every move of M by 2, move N by 1
	// N will point to the middle list
	for {
		if n.Next == nil {
			break
		}
		n = n.Next
		m = m.Next
		if m.Next == nil {
			break
		}
		m = m.Next
	}

	return n
}
