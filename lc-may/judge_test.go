package lcmay

import "testing"

func TestJudge1(t *testing.T) {
	var j int
	var a [][]int

	a = [][]int{
		{1, 2},
	}
	j = findJudge(2, a)
	if j != 2 {
		t.Error(j, " should be ", 2)
	}

	a = [][]int{
		{1, 3}, {2, 3},
	}

	j = findJudge(3, a)
	if j != 3 {
		t.Error(j, " should be ", 3)
	}

	a = [][]int{
		{1, 3}, {2, 3}, {3, 1},
	}

	j = findJudge(3, a)
	if j != -1 {
		t.Error(j, " should be ", -1)
	}

}
