package lc

import (
	"fmt"
	"sort"
	"strconv"
	"testing"
)

func countCases() [][]int {
	cases := make([][]int, 0)
	cases = append(cases, []int{1, 1, 2})
	cases = append(cases, []int{1, 1, 2, 2, 3, 3})
	cases = append(cases, []int{1, 2, 3})
	cases = append(cases, []int{1, 1, 3, 3, 5, 5, 7, 7})
	cases = append(cases, []int{1, 3, 2, 3, 5, 0})
	cases = append(cases, []int{1, 1, 2, 2})

	return cases
}

var countAnswers []int = []int{2, 4, 2, 0, 3, 2}

func TestCountElements(t *testing.T) {
	for i, v := range countCases() {
		c := countElements(v)
		if c != countAnswers[i] {
			t.Error(c, " should be ", countAnswers[i], i, v)
		}
	}
}

func qsort(a []int) {
	sort.IntSlice.Sort(a)
}

func qdataA(exp int) []int {
	a := make([]int, 0)
	v := 1
	for i := 0; i < exp; i++ {
		v *= 2
	}
	for i := 0; i < v; i++ {
		a = append(a, i)
	}
	return a
}

func qdataB(exp int) []int {
	a := make([]int, 0)
	v := 1
	for i := 0; i <= exp; i++ {
		v *= 2
	}
	for i := v; i >= 0; i-- {
		a = append(a, i)
	}
	return a
}

func BenchmarkQsort(b *testing.B) {
	var a []int
	var s string

	lo := 16
	hi := 20

	for exp := lo; exp <= hi; exp++ {
		a = qdataA(exp)
		s = strconv.FormatInt(int64(exp), 10)
		b.Run(s, func(b *testing.B) { qsort(a) })
	}
	fmt.Println()
	for exp := hi; exp >= lo; exp-- {
		a = qdataB(exp)
		s = strconv.FormatInt(int64(exp), 10)
		b.Run(s, func(b *testing.B) { qsort(a) })
	}
}
