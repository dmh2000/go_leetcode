package lc

import "fmt"

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// func newNode(v int) *TreeNode {
// 	return &TreeNode{Val: v, Left: nil, Right: nil}
// }

func insertLevelOrder(arr []int, root *TreeNode, i int, n int) *TreeNode {
	// Base case for recursion
	if i < n {
		if arr[i] == -2147483648 {
			return nil
		}

		root = newNode(arr[i])
		// root = temp

		//fmt.Println(root.Val)

		// insert left child
		//fmt.Println("left")
		root.Left = insertLevelOrder(arr, root.Left, 2*i+1, n)

		// insert right child
		//fmt.Println("right")
		root.Right = insertLevelOrder(arr, root.Right, 2*i+2, n)
	}
	return root
}

func inOrder(node *TreeNode) {
	if node != nil {
		inOrder(node.Left)
		fmt.Printf("%v  ", node.Val)
		inOrder(node.Right)
	}
}

func inOrder2(node *TreeNode, s string) {
	if node != nil {
		inOrder2(node.Left, "L")
		fmt.Printf("%v:%v  ", s, node.Val)
		inOrder2(node.Right, "R")
	}
}

func preOrder(node *TreeNode) {
	if node != nil {
		fmt.Printf("%v  ", node.Val)
		preOrder(node.Left)
		preOrder(node.Right)
	}
}

func preOrder2(node *TreeNode, s string) {
	if node != nil {
		fmt.Printf("%v:%v  ", s, node.Val)
		preOrder2(node.Left, "L")
		preOrder2(node.Right, "R")
	}
}

func postOrder(node *TreeNode) {
	if node != nil {
		postOrder(node.Left)
		postOrder(node.Right)
		fmt.Printf("%v  ", node.Val)
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

func newNode(v int) *TreeNode {
	return &TreeNode{Val: v, Left: nil, Right: nil}
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

func bstFromPreorder(preorder []int) *TreeNode {
	root := newNode(preorder[0])
	for i := 1; i < len(preorder); i++ {
		node := newNode(preorder[i])
		insertSorted(root, node)
	}
	return root
}
