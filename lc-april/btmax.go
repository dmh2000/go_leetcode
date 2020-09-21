package lc

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func bstMaxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func bstMax(node *TreeNode, sum *int) int {
	var left int
	var right int

	if node == nil {
		return 0
	}

	// recurse on the left side
	left = bstMax(node.Left, sum)

	// recurse on the right side
	right = bstMax(node.Right, sum)

	// maximum  for single path

	// either the left or right path is larger so pick the largest
	// chose left or right path, then add in the current node score
	// note : left or right path might be negative
	x := bstMaxInt(left, right) + node.Val

	// the resulting value might be negative so check it against
	// the current node score and take the largest
	x = bstMaxInt(x, node.Val)

	// now choose that value or the sum of left,right and node value
	y := bstMaxInt(x, left+right+node.Val)

	// update the sum with the larger of the current sum
	// or the new max
	*sum = bstMaxInt(*sum, y)

	// return the max of the left or right path + node value
	return x
}

func maxPathSum(root *TreeNode) int {
	sum := -2147483648
	bstMax(root, &sum)

	return sum
}
