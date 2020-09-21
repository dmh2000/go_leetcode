package lcjune

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(node *ListNode) {

	if node.Next == nil {
		return
	}

	for node.Next != nil {
		node.Val = node.Next.Val
		if node.Next.Next == nil {
			break
		}
		node = node.Next
	}
	node.Next = nil

}
