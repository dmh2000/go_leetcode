package lcmay

import "testing"

func TestRansom(t *testing.T) {

	b := canConstruct("a", "b")
	if b {
		t.Error(b, " should be ", false)
	}

	b = canConstruct("aa", "ab")
	if b {
		t.Error(b, " should be ", false)
	}

	b = canConstruct("aa", "aab")
	if !b {
		t.Error(b, " should be ", true)
	}

}
