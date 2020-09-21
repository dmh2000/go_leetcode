package lcmay

import (
	"fmt"
	"testing"
)

func TestKth1(t *testing.T) {
	var a []int
	var head *TreeNode
	var v int

	a = []int{3, 1, 4, TreeNull, 2}
	head = buildTree(a)
	head.inOrder()
	fmt.Println()

	v = kthSmallest(head, 1)
	if v != 1 {
		t.Error(v, " should be ", 1)
	}
	fmt.Println()

	a = []int{5, 3, 6, 2, 4, TreeNull, TreeNull, 1}
	head = buildTree(a)
	head.inOrder()
	fmt.Println()

	v = kthSmallest(head, 3)
	if v != 3 {
		t.Error(v, " should be ", 3)
	}
	fmt.Println()

	a = []int{5}
	head = buildTree(a)
	head.inOrder()
	fmt.Println()

	v = kthSmallest(head, 1)
	if v != 5 {
		t.Error(v, " should be ", 5)
	}
	fmt.Println()

}
