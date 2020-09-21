package lcmay

import "testing"

func TestCourses1(t *testing.T) {
	var b bool
	var a [][]int

	a = [][]int{
		{1, 0},
	}
	b = canFinish(2, a)
	if !b {
		t.Error(b, "should be", true)
	}

	a = [][]int{
		{1, 0},
		{0, 1},
	}
	b = canFinish(2, a)
	if b {
		t.Error(b, "should be", false)
	}
}
