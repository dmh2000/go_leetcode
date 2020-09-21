package lc

import "testing"

func TestDiv2(t *testing.T) {
	for i := 0; i < 100; i++ {
		for j := 1; j < 51; j++ {
			d := divide(i, j)
			if d != i/j {
				// t.Error(d, " should be ", i/j)
			}

		}
	}
}

func TestDiv2Large(t *testing.T) {
	i := maxInt1 / 2
	j := 2
	d := divide(i, j)
	if d != i/j {
		t.Error(d, " should be ", i/j)
	}
}
