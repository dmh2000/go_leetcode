package lc

import "fmt"

func parensGenStrings(s string, i int, p []string) []string {
	// base case
	if len(s) == 0 || i >= len(s) {
		// only even ones
		if len(s)%2 == 0 {
			p = append(p, s)
		}
		return p
	}

	if s[i] == '(' {
		// keep going
		p = parensGenStrings(s, i+1, p)
		// keep going
	} else if s[i] == ')' {
		p = parensGenStrings(s, i+1, p)
	} else if s[i] == '*' {
		// change * to (
		a := s[0:i] + "(" + s[i+1:]
		p = parensGenStrings(a, i+1, p)

		// change * to )
		b := s[0:i] + ")" + s[i+1:]
		p = parensGenStrings(b, i+1, p)

		// change * to ""
		c := s[0:i] + "" + s[i+1:]
		p = parensGenStrings(c, i+1, p)
	} else {
		panic("invalid character : " + string(s[i]))
	}

	return p
}

func parensValid(s string) bool {
	var b bool
	b = true
	count := 0
	for _, v := range s {
		if v == '(' {
			count++
		} else if v == ')' {
			count--
		}
		if count < 0 {
			b = false
		}
	}
	if count > 0 {
		b = false
	}
	return b
}

func checkValidString1(s string) bool {

	if len(s) == 0 {
		return true
	}

	// generate possible strings
	p := make([]string, 0)
	p = parensGenStrings(s, 0, p)

	// check for a valid one
	for _, v := range p {
		if parensValid(v) {
			return true
		}
	}

	return false
}

func checkValid2(count int, s string) int {
	if len(s) == 0 {
		return count
	}
	k := 0
	for i := range s {
		if s[i] == '(' {
			count++
		} else if s[i] == ')' {
			count--
			if count < 0 {
				break
			}
		} else if s[i] == '*' {
			k = checkValid2(count+1, s[i+1:])
			if k == 0 {
				count = k
				break
			}
			k = checkValid2(count-1, s[i+1:])
			if k == 0 {
				count = k
				break
			}
			k = checkValid2(count, s[i+1:])
			if k == 0 {
				count = k
				break
			}
			count += k
		}
	}
	return count
}

func checkValidString3(s string) bool {
	if len(s) == 0 {
		return true
	}

	b := true
	fwd := 0
	bak := 0
	n := len(s)
	for i := 0; i < n; i++ {
		// match parens going forward
		c := s[i]
		if c == '(' || c == '*' {
			fwd++
		} else {
			fwd--
		}

		// match parens going back
		c = s[n-i-1]
		if c == ')' || c == '*' {
			bak++
		} else {
			bak--
		}
		// if either count goes negative, its a fail
		if fwd < 0 || bak < 0 {
			b = false
		}
	}

	return b
}

func checkValidString(s string) bool {
	fmt.Println(s)
	b := checkValidString3(s)
	c := checkValid2(0, s)
	fmt.Println(s, b, c)
	return b
}
