package lcjune

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func invert(node *TreeNode) {
	if node == nil {
		return
	}
	var left *TreeNode
	left = node.Left
	node.Left = node.Right
	node.Right = left

	invert(node.Left)
	invert(node.Right)
}

func invertTree(root *TreeNode) *TreeNode {
	invert(root)
	return root
}
