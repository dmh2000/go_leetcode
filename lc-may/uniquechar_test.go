package lcmay

import "testing"

func TestUniqueChar1(t *testing.T) {

	i := firstUniqChar("leetcode")
	if i != 0 {
		t.Error(i, " should be ", 0)
	}

	i = firstUniqChar("loveleetcode")
	if i != 2 {
		t.Error(i, " should be ", 2)
	}

	i = firstUniqChar("eeeeeeaaaa")
	if i != -1 {
		t.Error(i, " should be ", -1)
	}

	i = firstUniqChar("xeeeeeeaaaa")
	if i != 0 {
		t.Error(i, " should be ", -1)
	}

	i = firstUniqChar("eeaax")
	if i != 4 {
		t.Error(i, " should be ", 4)
	}

}
