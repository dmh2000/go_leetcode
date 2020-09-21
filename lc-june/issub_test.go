package lcjune

import "testing"

func TestIsSub1(t *testing.T) {

	var s string
	var r string
	var b bool

	s = "abc"
	r = "ahbgdc"
	b = isSubsequence(s, r)
	if !b {
		t.Error(b, "should be", true)
	}

	s = "axc"
	r = "ahbgdc"
	b = isSubsequence(s, r)
	if b {
		t.Error(b, "should be", false)
	}
}
