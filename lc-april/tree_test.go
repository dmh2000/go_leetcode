package lc

import (
	"fmt"
	"testing"
)

func TestTree1(t *testing.T) {
	var root *TreeNode

	a := []int{1, 2, 3, 4, 5}

	root = insertLevelOrder(a, root, 0, len(a))

	inOrder(root)
}

func buildTree(a []int) *TreeNode {
	var root *TreeNode

	root = insertLevelOrder(a, root, 0, len(a))

	return root
}

func TestDiameter(t *testing.T) {
	var a []int
	var d int
	var root *TreeNode

	a = []int{1, 2, 3, 4, 5}
	root = buildTree(a)
	d = diameterOfBinaryTree(root)
	if d != 3 {
		t.Error(d, " should be ", 3)
	}

	a = []int{1, 2}
	root = buildTree(a)
	d = diameterOfBinaryTree(root)
	if d != 1 {
		t.Error(d, " should be ", 1)
	}

	a = []int{2, 3, -2147483648, 1}
	root = buildTree(a)
	d = diameterOfBinaryTree(root)
	if d != 2 {
		t.Error(d, " should be ", 2)
	}
}

func TestPreorder(t *testing.T) {
	a := []int{8, 5, 1, 7, 10, 12}
	root := bstFromPreorder(a)
	preOrder(root)
}

func TestBtmax1(t *testing.T) {

	a := []int{1, 2, 3}
	root := buildTree(a)
	x := maxPathSum(root)
	if x != 6 {
		t.Error(x, " should be ", 6)
	}
	fmt.Println()

	a = []int{-10, 9, 20, -2147483648, -2147483648, 15, 7}
	root = buildTree(a)
	x = maxPathSum(root)
	if x != 42 {
		t.Error(x, " should be ", 42)
	}
	fmt.Println()

	a = []int{-3}
	root = buildTree(a)
	x = maxPathSum(root)
	if x != -3 {
		t.Error(x, " should be ", -3)
	}
	fmt.Println()

	a = []int{2, -1}
	root = buildTree(a)
	x = maxPathSum(root)
	if x != 2 {
		t.Error(x, " should be ", 2)
	}
	fmt.Println()

	a = []int{1, -2, 3}
	root = buildTree(a)
	x = maxPathSum(root)
	if x != 4 {
		t.Error(x, " should be ", 4)
	}
	fmt.Println()
}

func TestBtmax2(t *testing.T) {
	// 5,4,8,11,null,13,4,7,2,null,null,null,1

	a := []int{5, 4, 8, 11, -2147483648, 13, 4, 7, 2, -2147483648, -2147483648, -2147483648, 1}
	root := buildTree(a)
	preOrder(root)
	fmt.Println()
	x := maxPathSum(root)
	if x != 48 {
		t.Error(x, " should be ", 48)
	}
	fmt.Println()
}

func TestValidString1(t *testing.T) {
	a := []int{0, 1, 0, 0, 1, 0, -2147483648, -2147483648, 1, 0, 0}
	root := buildTree(a)
	// preOrder2(root, "T")
	// fmt.Println()
	// inOrder2(root, "T")
	arr := []int{0, 1, 0, 1}
	b := isValidSequence(root, arr)
	if !b {
		t.Error(b, " should be ", true)
	}

	a = []int{0, 1, 0, 0, 1, 0, -2147483648, -2147483648, 1, 0, 0}
	root = buildTree(a)
	arr = []int{0, 0, 1}
	b = isValidSequence(root, arr)
	if b {
		t.Error(b, " should be ", false)
	}

	a = []int{0, 1, 0, 0, 1, 0, -2147483648, -2147483648, 1, 0, 0}
	root = buildTree(a)
	arr = []int{0, 1, 1}
	b = isValidSequence(root, arr)
	if b {
		t.Error(b, " should be ", false)
	}
}
