package lcmay

import "container/list"

func inOrderKrecursive(node *TreeNode, i *int, k int, v *int) bool {
	var done bool
	if node != nil {
		done = inOrderKrecursive(node.Left, i, k, v)
		if done {
			return true
		}

		// Visit
		*i++
		// fmt.Println(node.Val, *i, k)
		if *i == k {
			*v = node.Val
			return true
		}

		done = inOrderKrecursive(node.Right, i, k, v)
		if done {
			return true
		}
	}
	return false
}

func inOrderKstack(node *TreeNode, i *int, k int, v *int) {
	var stack list.List
	var last *list.Element
	for {
		for node != nil {
			stack.PushFront(node)
			node = node.Left
		}

		if stack.Len() == 0 {
			break
		}

		// pop the current value
		last = stack.Front()
		stack.Remove(last)
		node = last.Value.(*TreeNode)

		// Visit
		*i++
		// fmt.Println(node.Val, *i, k)
		if *i == k {
			*v = node.Val
			return
		}

		// right branch
		node = last.Value.(*TreeNode).Right
	}

}

func kthSmallest(root *TreeNode, k int) int {
	var v int
	var i int

	i = 0
	inOrderKstack(root, &i, k, &v)

	return v
}
