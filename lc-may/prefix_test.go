package lcmay

import "testing"

func TestLCP(t *testing.T) {
	a := []string{
		"flower",
		"flow",
		"flight",
	}

	p := longestCommonPrefix(a)
	if p != "fl" {
		t.Error(p, " should be ", "fl")
	}

	a = []string{
		"dog",
		"racecar",
		"car",
	}

	p = longestCommonPrefix(a)
	if p != "" {
		t.Error(p, " should be empty string")
	}

	a = []string{
		"dog",
		"dick",
		"done",
	}

	p = longestCommonPrefix(a)
	if p != "d" {
		t.Error(p, " should be ", "d")
	}

	a = []string{}

	p = longestCommonPrefix(a)
	if p != "" {
		t.Error(p, " should be ", "")
	}

	a = []string{
		"a",
	}

	p = longestCommonPrefix(a)
	if p != "a" {
		t.Error(p, " should be ", "")
	}
	a = []string{
		"a",
		"a",
	}

	p = longestCommonPrefix(a)
	if p != "a" {
		t.Error(p, " should be ", "a")
	}

	a = []string{
		"a",
		"b",
	}

	p = longestCommonPrefix(a)
	if p != "" {
		t.Error(p, " should be ", "empty")
	}
}
