package lc

import (
	"fmt"
	"testing"
)

func TestMaxSquare(t *testing.T) {
	m := [][]byte{}

	sq := maximalSquare(m)
	if sq != 0 {
		t.Error(sq, " should equal ", 0)
	}

	m = [][]byte{
		{'1'},
	}
	sq = maximalSquare(m)
	if sq != 1 {
		t.Error(sq, " should equal ", 1)
	}

	m = [][]byte{
		{'0'},
	}
	sq = maximalSquare(m)
	if sq != 0 {
		t.Error(sq, " should equal ", 0)
	}

	m = [][]byte{
		{'0', '1', '0'},
	}
	sq = maximalSquare(m)
	if sq != 1 {
		t.Error(sq, " should equal ", 1)
	}

	m = [][]byte{
		{'0'},
		{'1'},
	}
	sq = maximalSquare(m)
	if sq != 1 {
		t.Error(sq, " should equal ", 1)
	}

	m = [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}

	sq = maximalSquare(m)
	if sq != 4 {
		t.Error(sq, " should equal ", 4)
	}

	fmt.Println()
	m = [][]byte{
		{'1', '0', '1', '1', '1'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}

	sq = maximalSquare(m)
	if sq != 9 {
		t.Error(sq, " should equal ", 9)
	}

	fmt.Println()
	m = [][]byte{
		{'1', '1', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
	}

	sq = maximalSquare(m)
	if sq != 16 {
		t.Error(sq, " should equal ", 16)
	}
}
