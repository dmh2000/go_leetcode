package lcmay

import (
	"fmt"
	"testing"
)

type rk struct {
	num string
	k   int
	out string
}

var rkval []rk = []rk{
	{"11119999", 2, "111199"},
	{"54321", 3, "21"},
	{"12345", 3, "12"},
	{"1432219", 3, "1219"},
	{"10200", 1, "200"},
	{"10", 2, "0"},
	{"112", 1, "11"},
}

func TestRemoveK1(t *testing.T) {
	var s string

	for i := 0; i < len(rkval); i++ {
		s = removeKdigits(rkval[i].num, rkval[i].k)
		fmt.Println(s)
		if s != rkval[i].out {
			t.Error(s, " should be ", rkval[i].out)
		}
	}
}
