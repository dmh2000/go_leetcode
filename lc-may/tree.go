package lcmay

import "fmt"

// TreeNull ...
const TreeNull = -2147483648

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertLevelOrder(arr []int, root *TreeNode, i int, n int) *TreeNode {
	// Base case for recursion
	if i >= n {
		return root
	}

	if arr[i] == TreeNull {
		return nil
	}

	root = newNode(arr[i])

	// insert left child
	root.Left = insertLevelOrder(arr, root.Left, 2*i+1, n)

	// insert right child
	root.Right = insertLevelOrder(arr, root.Right, 2*i+2, n)

	return root
}

func inOrderF(node *TreeNode, f func(n *TreeNode)) {
	if node != nil {
		inOrder(node.Left)
		// Visit
		f(node)
		inOrder(node.Right)
	}
}

func preOrderF(node *TreeNode, f func(n *TreeNode)) {
	if node != nil {
		// Visit
		f(node)
		preOrder(node.Left)
		preOrder(node.Right)
	}
}

func postOrderF(node *TreeNode, f func(n *TreeNode)) {
	if node != nil {
		postOrder(node.Left)
		postOrder(node.Right)
		// Visit
		f(node)
	}
}

func newNode(v int) *TreeNode {
	return &TreeNode{Val: v, Left: nil, Right: nil}
}

func buildLeetCodeTree(a []int) *TreeNode {
	var root *TreeNode

	root = insertLevelOrder(a, root, 0, len(a))

	return root
}

// ====================================================

func inOrder(node *TreeNode) {
	if node != nil {
		inOrder(node.Left)
		// Visit
		fmt.Print(node.Val, ",")
		inOrder(node.Right)
	}
}

func (node *TreeNode) inOrder() {
	if node != nil {
		node.Left.inOrder()
		// Visit
		fmt.Print(node.Val, ",")
		node.Right.inOrder()
	}
}

func preOrder(node *TreeNode) {
	if node != nil {
		// Visit
		fmt.Print(node.Val, ",")
		preOrder(node.Left)
		preOrder(node.Right)
	}
}

func (node *TreeNode) preOrder() {
	if node != nil {
		// Visit
		fmt.Print(node.Val, ",")
		node.Left.preOrder()
		node.Right.preOrder()
	}
}

func postOrder(node *TreeNode) {
	if node != nil {
		postOrder(node.Left)
		postOrder(node.Right)
		// Visit
		fmt.Print(node.Val, ",")
	}
}

func (node *TreeNode) postOrder() {
	if node != nil {
		node.Left.postOrder()
		node.Right.postOrder()
		// Visit
		fmt.Print(node.Val, ",")
	}
}

func buildTree(a []int) *TreeNode {
	var root *TreeNode

	root = insertLevelOrder(a, root, 0, len(a))

	return root
}

func preOrder2(node *TreeNode, s string) {
	if node != nil {
		// Visit
		preOrder2(node.Left, "L")
		preOrder2(node.Right, "R")
	}
}

func inOrderH(node *TreeNode, h interface{}) {
	if node != nil {
		inOrderH(node.Left, "L")
		// fmt.Printf("%v:%v  ", h, node.Val)
		inOrderH(node.Right, "R")
	}
}

func bstFromPreorder(preorder []int) *TreeNode {
	root := newNode(preorder[0])
	for i := 1; i < len(preorder); i++ {
		node := newNode(preorder[i])
		insertSorted(root, node)
	}
	return root
}

func diameterOfBinaryTree(root *TreeNode) int {
	max := 0
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 0
	}
	findLongest(root, &max)

	return max
}

func insertSorted(root *TreeNode, node *TreeNode) {
	if node.Val < root.Val {
		if root.Left == nil {
			root.Left = node
			return
		}
		insertSorted(root.Left, node)
	} else {
		if root.Right == nil {
			root.Right = node
			return
		}
		insertSorted(root.Right, node)
	}
}

func findLongest(node *TreeNode, max *int) int {
	var r int
	var left int
	var right int

	if node == nil {
		return 0
	}

	left = 0
	if node.Left != nil {
		left = findLongest(node.Left, max)
	}

	fmt.Printf("%v  ", node.Val)

	right = 0
	if node.Right != nil {
		right = findLongest(node.Right, max)
	}

	r = left + right
	if r > *max {
		*max = r
	}

	if left > right {
		r = left
	} else {
		r = right
	}

	return r + 1
}
