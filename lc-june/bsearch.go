package lcjune

// TreeNode Definition for a binary tree node.
// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func preOrderSearch(node *TreeNode, val int) *TreeNode {
	var p *TreeNode

	if node == nil {
		return nil
	}

	if node.Val == val {
		return node
	}

	p = preOrderSearch(node.Left, val)
	if p != nil {
		return p
	}
	p = preOrderSearch(node.Right, val)
	if p != nil {
		return p
	}

	return nil
}
func searchBST(root *TreeNode, val int) *TreeNode {
	return preOrderSearch(root, val)
}
