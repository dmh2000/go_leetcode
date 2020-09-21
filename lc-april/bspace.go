package lc

import (
	"bytes"
	"fmt"
)

func bspace1(s string) string {
	var b bytes.Buffer
	var state int

	state = 0
	for _, v := range s {
		switch state {
		case 0:
			// start state
			// looking for character or backspace
			if v == rune('#') {
				// skip
			} else {
				b.WriteRune(v)
			}
			state = 1
		case 1:
			if v == rune('#') {
				if b.Len() > 0 {
					b.Truncate(b.Len() - 1)
				}
			} else {
				b.WriteRune(v)
			}
		}
	}
	return b.String()
}

func bspace(s string) string {
	i := 0
	for i < len(s) {
		if len(s) == 0 {
			s = ""
			break
		} else if s[i] == '#' {
			if len(s) == 1 {
				s = ""
				break
			} else if i == 0 {
				s = s[1:]
			} else {
				s = s[0:i-1] + s[i+1:]
				i--
			}
		} else {
			i++
		}
	}
	return s
}

func backspaceCompare(S string, T string) bool {

	s := bspace(S)
	t := bspace(T)
	fmt.Println(s, t)
	return s == t
}
