package lcjune

import (
	"fmt"
	"testing"
)

func TestInvert1(t *testing.T) {

	var root *TreeNode

	root = buildLeetCodeTree([]int{1, 2, 3, 4, 5, 6, 7})
	preOrder2(root, "root", 0)
	invertTree(root)
	fmt.Println()
	preOrder2(root, "root", 0)

	fmt.Println()

	root = buildLeetCodeTree([]int{1, 2, 3, 4, 5, 6})
	preOrder2(root, "root", 0)
	invertTree(root)
	fmt.Println()
	preOrder2(root, "root", 0)
}

func TestInvert2(t *testing.T) {

	var root *TreeNode

	root = bstFromPreorder([]int{1, 2, 3, 4, 5, 6, 7})
	inOrder(root)
	invertTree(root)
	fmt.Println()
	inOrder(root)

}
