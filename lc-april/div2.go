package lc

import "fmt"

const maxInt1 = 2147483647
const minInt1 = -2147483648

// negate without multiply
func negate(v int) int {
	v = ^v
	return v + 1
}

// absolute without multiply
func abs(v int) int {
	if v < 0 {
		v = negate(v)
	}
	return v
}

// I had to look this one up. i knew there would be a method
// using bit shifting but couldn't figure it out
func divide(dividend int, divisor int) int {
	var sign int
	var quotient int
	var temp int
	var one int

    // get the resulting sign of the two terms
	sign = 1
	if dividend < 0 && divisor > 0 {
		sign = -1
	} else if dividend > 0 && divisor < 0 {
		sign = -1
	}

    // make them both positive
	dividend = abs(dividend)
	divisor = abs(divisor)

	quotient = 0
	temp = 0
	one = 1
    // scan down to align divisor with dividend
	for i := 31; i >= 0; i-- {
		d := divisor << i
        // here's the tricky part
		// if the current divisor + the divisor shifted right one bit is less than the dividend
		fmt.Printf("%33b : %d\n",dividend,dividend)
		fmt.Printf("%33b : %d\n",d,d)
		fmt.Printf("%33b : %d\n",temp,temp)
		if (temp + d) <= dividend {
            // remember the current divisor
			temp += d
            // set a bit for this term
			quotient |= one << i
		}
		fmt.Printf("%33b : %d\n",quotient,quotient)
		fmt.Println()
	}

    // set the correct sign
	if sign < 0 {
		quotient = negate(quotient)
	}

	// adjust for golang 64 bit integers that don't overflow 
    // the way this problem wants it to
	if quotient == (maxInt1 + 1) {
		quotient--
	}

    // check for other overflows
	if quotient > maxInt1 {
		quotient = -1
	} else if quotient < minInt1 {
		quotient = -1
	}

	return quotient
}