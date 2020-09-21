package lc

import (
	"testing"
)

func TestBackspace(t *testing.T) {
	var b bool
	b = backspaceCompare("ab#c", "ad#c")
	if !b {
		t.Error(b, " should be ", true)
	}

	b = backspaceCompare("ab##", "ad##")
	if !b {
		t.Error(b, " should be ", true)
	}

	b = backspaceCompare("a##c", "#a#c")
	if !b {
		t.Error(b, " should be ", true)
	}

	b = backspaceCompare("a#c", "b")
	if b {
		t.Error(b, " should be ", false)
	}

}
