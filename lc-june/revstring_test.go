package lcjune

import (
	"testing"
)

func TestRev1(t *testing.T) {
	var s string
	var r string
	var b []byte

	s = "123456789"
	r = "987654321"
	b = []byte(s)
	reverseString(b)
	if string(b) != r {
		t.Error(s, "should be", r)
	}

	s = "12345678"
	r = "87654321"
	b = []byte(s)
	reverseString(b)
	if string(b) != r {
		t.Error(s, "should be", r)
	}

	s = ""
	r = ""
	b = []byte(s)
	reverseString(b)
	if string(b) != r {
		t.Error(s, "should be", r)
	}

	s = "1"
	r = "1"
	b = []byte(s)
	reverseString(b)
	if string(b) != r {
		t.Error(s, "should be", r)
	}

	s = "12"
	r = "21"
	b = []byte(s)
	reverseString(b)
	if string(b) != r {
		t.Error(s, "should be", r)
	}

}
