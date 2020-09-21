package lc

import (
	"fmt"
	"testing"
)

func TestLcs1(t *testing.T) {
	a := "abzcde"
	b := "acxxe"

	fmt.Println(a, b)
	v := longestCommonSubsequence(a, b)
	if v != 3 {
		t.Error(v, " should be ", 3)
	}

	a = "abczdefg"
	b = "zazceq"

	v = longestCommonSubsequence(a, b)
	if v != 3 {
		t.Error(v, " should be ", 3)
	}
}

func TestLcs2(t *testing.T) {
	a := ""
	b := ""
	for i := 0; i < 900; i++ {
		a += "y"
		b += "z"
	}
	a = a + "ac"
	b = "a" + b + "c"

	v := longestCommonSubsequence(a, b)
	if v != 2 {
		t.Error(v, " should be ", 2)
	}
}

func TestLcs3(t *testing.T) {
	a := ""
	b := ""
	for i := 0; i < 900; i++ {
		a += "y"
		b += "z"
	}

	v := longestCommonSubsequence(a, b)
	if v != 0 {
		t.Error(v, " should be ", 0)
	}
}
