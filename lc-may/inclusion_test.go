package lcmay

import (
	"testing"
)

func TestInclusion1(t *testing.T) {
	var b bool

	b = checkInclusion("ba", "eidabooo")
	if !b {
		t.Error(b, " should be ", true)
	}

	b = checkInclusion("abc", "ccccbbbbaaaa")
	if b {
		t.Error(b, " should be ", false)
	}

	b = checkInclusion("adc", "dcda")
	if !b {
		t.Error(b, " should be ", true)
	}

	b = checkInclusion("ab", "a")
	if b {
		t.Error(b, " should be ", false)
	}

	b = checkInclusion("a", "a")
	if !b {
		t.Error(b, " should be ", true)
	}

}

func TestInclusion2(t *testing.T) {
	var b bool
	var s1 string
	var s2 string

	s1 = "cab"

	s2 = "lgkjkgjkgjkjgkjgkjgkjgg"
	s2 += "abc"
	s2 += "lgkjkgjkgjkjgkjgkjgkjgg"
	b = checkInclusion(s1, s2)
	if !b {
		t.Error(b, " should be ", true)
	}
}

func TestInclusion3(t *testing.T) {
	var b bool
	var s1 string
	var s2 string

	s1 = "lgkjkgjkgjkjgkjgkjgkjgg"

	s := "lasldmvididuyaaiduwheodlaksj"
	for i := 0; i < 1; i++ {
		s2 += s
		s2 += "glkjkgjkgjkjgkjgkjgk"
	}
	s2 += "z"
	s2 += s1

	for i := 0; i < 1; i++ {
		s2 += s
	}

	b = checkInclusion(s1, s2)
	if !b {
		t.Error(b, " should be ", true)
	}
}

func TestInclusionAll(t *testing.T) {
	t.Run("1", TestInclusion1)
	t.Run("2", TestInclusion2)
	t.Run("3", TestInclusion3)
}
