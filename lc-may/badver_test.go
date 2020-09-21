package lcmay

import "testing"

func TestBadVersion(t *testing.T) {
	for n := 1; n < 10; n++ {
		setNN(n)
		b := firstBadVersion(n)
		if b != n {
			t.Error(b, " should be ", n)
		}
	}
}
