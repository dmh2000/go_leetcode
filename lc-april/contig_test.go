package lc

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestContig(t *testing.T) {
	var m int

	m = findMaxLength([]int{0, 1, 1, 0, 1, 1, 1, 0})
	if m != 4 {
		t.Error(m, "should be ", 4)
	}

	m = findMaxLength([]int{0, 1})
	if m != 2 {
		t.Error(m, "should be ", 2)
	}

	m = findMaxLength([]int{0, 1, 0})
	if m != 2 {
		t.Error(m, "should be ", 2)
	}

	m = findMaxLength([]int{1, 1, 1})
	if m != 0 {
		t.Error(m, "should be ", 0)
	}

	m = findMaxLength([]int{0, 0, 0})
	if m != 0 {
		t.Error(m, "should be ", 0)
	}

	m = findMaxLength([]int{0, 0, 0, 1})
	if m != 2 {
		t.Error(m, "should be ", 2)
	}
}

func BenchmarkContigB1(b *testing.B) {
	a := make([]int, 1000)
	c := 1
	for i := 0; i < len(a); i++ {
		a[i] = c
		if c == 1 {
			c = 0
		} else {
			c = 1
		}
	}
	findMaxLength(a)
}

func BenchmarkContigB2(b *testing.B) {
	a := make([]int, 50000)
	c := 1
	for i := 0; i < len(a); i++ {
		a[i] = c
		if c == 1 {
			c = 0
		} else {
			c = 1
		}
	}
	fmt.Println(findMaxLength(a))
}

func BenchmarkContigB3(b *testing.B) {
	a := make([]int, 50000)
	for i := 0; i < len(a); i++ {
		a[i] = rand.Int() % 2
	}
	findMaxLength(a)
}
