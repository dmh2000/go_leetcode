package lcmay

import "testing"

func TestJewels(t *testing.T) {
	j := "aA"
	s := "aAAbbbb"

	n := numJewelsInStones1(j, s)
	if n != 3 {
		t.Error(n, " should be ", 3)
	}

	n = numJewelsInStones2(j, s)
	if n != 3 {
		t.Error(n, " should be ", 3)
	}

	j = "z"
	s = "ZZ"

	n = numJewelsInStones1(j, s)
	if n != 0 {
		t.Error(n, " should be ", 0)
	}

	n = numJewelsInStones2(j, s)
	if n != 0 {
		t.Error(n, " should be ", 0)
	}
}
