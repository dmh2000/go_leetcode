package lcmay

import "testing"

func TestEdit1(t *testing.T) {
	var k int

	k = minDistance("horse", "ros")
	if k != 3 {
		t.Error(k, "should be", 3)
	}

	k = minDistance("intention", "execution")
	if k != 5 {
		t.Error(k, "should be", 5)
	}
}
