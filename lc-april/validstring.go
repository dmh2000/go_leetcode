package lc

func valSeq(node *TreeNode, arr []int, i int) bool {
	if node == nil {
		// should only happen if root is nil
		return false
	}

	// if the input arr is exhausted but there are
	// more nodes to test, its false
	if i >= len(arr) {
		return false
	}

	// no match on this node
	if node.Val != arr[i] {
		return false
	}

	// if at end of input array, and there are no child nodes
	// it is true
	if node.Left == nil && node.Right == nil && i == len(arr)-1 {
		return true
	}

	lx := false
	// check left node
	if node.Left != nil {
		lx = valSeq(node.Left, arr, i+1)
		if lx {
			return true
		}
	}

	// check right node
	rx := false
	if node.Right != nil {
		rx = valSeq(node.Right, arr, i+1)
		if rx {
			return true
		}
	}

	// return result
	return false
}

func isValidSequence(root *TreeNode, arr []int) bool {
	b := valSeq(root, arr, 0)
	return b
}
