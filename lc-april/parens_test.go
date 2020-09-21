package lc

import (
	"testing"
)

func TestParens1(t *testing.T) {
	var b bool

	b = checkValidString("(())((())()()(*)(*()(())())())()()((()())((()))(*")
	if b {
		t.Error(b, " should be ", false)
	}

	// true
	b = checkValidString("(*()")
	if !b {
		t.Error(b, " should be ", true)
	}
	b = checkValidString("*")
	if !b {
		t.Error(b, " should be ", true)
	}
	b = checkValidString("**")
	if !b {
		t.Error(b, " should be ", true)
	}
	b = checkValidString("(*)")
	if !b {
		t.Error(b, " should be ", true)
	}
	b = checkValidString("(**)()()()")
	if !b {
		t.Error(b, " should be ", true)
	}
	b = checkValidString("(*))")
	if !b {
		t.Error(b, " should be ", true)
	}

	// false
	b = checkValidString("(")
	if b {
		t.Error(b, " should be ", false)
	}
	b = checkValidString("*))")
	if b {
		t.Error(b, " should be ", false)
	}
	b = checkValidString("(())(")
	if b {
		t.Error(b, " should be ", false)
	}

	//b = checkValidString("(())()())(*))(((((())()*))**))**()(*(()()*)(**(())()**)((**(()(((()()**)())*(*))(())()()*")
	//fmt.Println(b)

}
