package lcmay

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type cousins struct {
	x       int
	xlevel  int
	xparent *TreeNode

	y       int
	ylevel  int
	yparent *TreeNode
}

func preOrderCousins(node *TreeNode, depth int, parent *TreeNode, v *cousins) {

	if node != nil {
		if node.Val == v.x {
			v.xlevel = depth
			v.xparent = parent
		} else if node.Val == v.y {
			v.ylevel = depth
			v.yparent = parent
		}

		preOrderCousins(node.Left, depth+1, node, v)

		preOrderCousins(node.Right, depth+1, node, v)
	}
}

func isCousins(root *TreeNode, x int, y int) bool {
	c := cousins{x, 0, nil, y, 0, nil}

	preOrderCousins(root, 0, nil, &c)

	return (c.ylevel == c.xlevel) && (c.yparent != c.xparent)
}
