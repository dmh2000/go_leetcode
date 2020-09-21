package lc

import "testing"

var rws1 = "  a good   example "
var rwa1 = "example good a"

func TestRevWords(t *testing.T) {
	s := reverseWords(rws1)
	if s != rwa1 {
		t.Error(s, "should be", rwa1)
	}
}
