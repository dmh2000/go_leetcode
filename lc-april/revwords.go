package lc

import (
	"bytes"
	"strings"
)

func reverseWords1(s string) string {
	s = strings.TrimSpace(s)
	a := strings.Split(s, " ")
	b := make([]string, 0)
	for _, v := range a {
		if v != "" {
			b = append([]string{v}, b...)
		}
	}
	r := strings.Join(b, " ")
	return r
}

func reverseWords(s string) string {
	var state int
	var a []string
	var b bytes.Buffer

	a = make([]string, 0)
	state = 0
	for _, v := range s {
		c := string(v)
		switch state {
		case 0:
			if c == " " {
				// leading space, skip it
				b.Reset()
			} else {
				// first character in a word
				b.WriteString(c)
				state = 1
			}
		case 1:
			// next character in first word
			if c != " " {
				b.WriteString(c)
			} else {
				// intermediate space, keep the word
				a = append(a, b.String())
				b.Reset()
				state = 2
			}
		case 2:
			// looking for additional spaces or next word
			if c == " " {
				// additional space, skip it
			} else {
				// another word
				b.WriteString(c) // start next word
				state = 3
			}
		case 3:
			// next character
			if c != " " {
				b.WriteString(c)
			} else {
				// intermediate space, keep the word
				a = append(a, b.String())
				b.Reset()
				state = 2
			}
		}
	}
	// remaining word
	if b.Len() > 0 {
		a = append(a, b.String())
	}

	b.Reset()
	for i := len(a) - 1; i >= 1; i-- {
		b.WriteString(a[i])
		b.WriteString(" ")
	}
	b.WriteString(a[0])

	return b.String()
}
