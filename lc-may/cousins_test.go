package lcmay

import (
	"fmt"
	"testing"
)

func TestCousins(t *testing.T) {
	var a []int
	var root *TreeNode
	var b bool

	a = []int{1, 2, 3, 4}
	root = buildTree(a)

	fmt.Println("-------------------------------------")
	b = isCousins(root, 4, 3)
	if b {
		t.Error(b, " should be ", false)
	}

	a = []int{1, 2, 2, 4}
	root = buildTree(a)

	fmt.Println("-------------------------------------")
	b = isCousins(root, 2, 2)
	if b {
		t.Error(b, " should be ", false)
	}

	a = []int{1, 2, 3, -2147483648, 4, -2147483648, 5}
	root = buildTree(a)

	fmt.Println("-------------------------------------")
	b = isCousins(root, 5, 4)
	if !b {
		t.Error(b, " should be ", true)
	}

	// 	a = []int{1, 2, 3, -2147483648, 4}
	// 	root = buildTree(a)

	// 	fmt.Println("-------------------------------------")
	// 	b = isCousins(root, 2, 3)
	// 	if !b {
	// 		t.Error(b, " should be ", true)
	// 	}
}
