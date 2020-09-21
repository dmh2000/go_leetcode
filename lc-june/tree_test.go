package lcjune

import (
	"fmt"
	"testing"
)

func visit(node *TreeNode) {
	fmt.Print(node.Val, ",")
}

func TestTree1(t *testing.T) {
	root := buildLeetCodeTree([]int{1, 2, 3, 4, 5})
	preOrderF(root, visit)
	fmt.Println()
}

func TestTree2(t *testing.T) {
	root := buildLeetCodeTree([]int{1, 2, 3, 4, 5})
	inOrderF(root, visit)
	fmt.Println()
}

func TestTree3(t *testing.T) {
	root := buildLeetCodeTree([]int{1, 2, 3, 4, 5})
	postOrderF(root, visit)
	fmt.Println()
}

func TestTreeAll(t *testing.T) {
	t.Run("preorder", TestTree1)
	t.Run("inorder", TestTree2)
	t.Run("postorder", TestTree3)
}

func TestPreorder(t *testing.T) {
	a := []int{8, 5, 1, 7, 10, 12}
	root := bstFromPreorder(a)
	preOrder(root)
}
