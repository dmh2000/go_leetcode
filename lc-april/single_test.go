package lc

import (
	"flag"
	"fmt"
	"os"
	"testing"
)

var count = 1024

func singleNumberT(f func([]int) int, t *testing.T) {
	// lots of numbers
	w := count - 1
	a := make([]int, 0)
	for i := 0; i < w; i++ {
		a = append(a, i)
	}
	for i := 0; i < w; i++ {
		a = append(a, i)
	}
	a = append(a, count)

	b := f(a)

	if b != count {
		t.Error(b, " should be ", count)
	}
	// increment for next test
	count *= 2
}

func singleNumberB(f func([]int) int, b *testing.B) {
	// lots of numbers
	w := count - 1
	a := make([]int, 0)
	for i := 0; i < w; i++ {
		a = append(a, i)
	}
	for i := 0; i < w; i++ {
		a = append(a, i)
	}
	a = append(a, count)

	x := f(a)

	if x != count {
		b.Error(x, " should be ", count)
	}
	// increment for next test
	count *= 2

}

func TestSingleNumber(t *testing.T) {
	f := singleNumberHeap

	t.Run(fmt.Sprintf("%v", count), func(t *testing.T) { singleNumberT(f, t) })
	t.Run(fmt.Sprintf("%v", count), func(t *testing.T) { singleNumberT(f, t) })
	t.Run(fmt.Sprintf("%v", count), func(t *testing.T) { singleNumberT(f, t) })
	t.Run(fmt.Sprintf("%v", count), func(t *testing.T) { singleNumberT(f, t) })
	t.Run(fmt.Sprintf("%v", count), func(t *testing.T) { singleNumberT(f, t) })
	// t.Run(fmt.Sprintf("%v", count), func(t *testing.T) { singleNumberT(f, t) })
	// t.Run(fmt.Sprintf("%v", count), func(t *testing.T) { singleNumberT(f, t) })
	// t.Run(fmt.Sprintf("%v", count), func(t *testing.T) { singleNumberH(f, t) })
	// t.Run(fmt.Sprintf("%v", count), func(t *testing.T) { singleNumberH(f, t) })
	// t.Run(fmt.Sprintf("%v", count), func(t *testing.T) { singleNumberH(f, t) })
}

func TestSingleNumberM(t *testing.T) {
	f := singleNumberMap
	count = 1024

	t.Run(fmt.Sprintf("%v", count), func(t *testing.T) { singleNumberT(f, t) })
	t.Run(fmt.Sprintf("%v", count), func(t *testing.T) { singleNumberT(f, t) })
	t.Run(fmt.Sprintf("%v", count), func(t *testing.T) { singleNumberT(f, t) })
	t.Run(fmt.Sprintf("%v", count), func(t *testing.T) { singleNumberT(f, t) })
	t.Run(fmt.Sprintf("%v", count), func(t *testing.T) { singleNumberT(f, t) })
	t.Run(fmt.Sprintf("%v", count), func(t *testing.T) { singleNumberT(f, t) })
	t.Run(fmt.Sprintf("%v", count), func(t *testing.T) { singleNumberT(f, t) })
	// t.Run(fmt.Sprintf("%v", count), func(t *testing.T) { singleNumberH(f, t) })
	// t.Run(fmt.Sprintf("%v", count), func(t *testing.T) { singleNumberH(f, t) })
	// t.Run(fmt.Sprintf("%v", count), func(t *testing.T) { singleNumberH(f, t) })
}

func TestSingleNumberM2(t *testing.T) {
	f := singleNumberMap2
	count = 1024
	t.Run(fmt.Sprintf("%v", count), func(t *testing.T) { singleNumberT(f, t) })
	t.Run(fmt.Sprintf("%v", count), func(t *testing.T) { singleNumberT(f, t) })
	t.Run(fmt.Sprintf("%v", count), func(t *testing.T) { singleNumberT(f, t) })
	t.Run(fmt.Sprintf("%v", count), func(t *testing.T) { singleNumberT(f, t) })
	t.Run(fmt.Sprintf("%v", count), func(t *testing.T) { singleNumberT(f, t) })
	t.Run(fmt.Sprintf("%v", count), func(t *testing.T) { singleNumberT(f, t) })
	t.Run(fmt.Sprintf("%v", count), func(t *testing.T) { singleNumberT(f, t) })
	// t.Run(fmt.Sprintf("%v", count), func(t *testing.T) { singleNumberH(f, t) })
	// t.Run(fmt.Sprintf("%v", count), func(t *testing.T) { singleNumberH(f, t) })
	// t.Run(fmt.Sprintf("%v", count), func(t *testing.T) { singleNumberH(f, t) })
}
func BenchmarkSingleNumber(b *testing.B) {
	f := singleNumberHeap

	b.Run(fmt.Sprintf("%v", count), func(b *testing.B) { singleNumberB(f, b) })
	b.Run(fmt.Sprintf("%v", count), func(b *testing.B) { singleNumberB(f, b) })
	b.Run(fmt.Sprintf("%v", count), func(b *testing.B) { singleNumberB(f, b) })
	// b.Run(fmt.Sprintf("%v", count), func(b *testing.B) { singleNumberB(f, b) })
	// b.Run(fmt.Sprintf("%v", count), func(b *testing.B) { singleNumberB(f, b) })
	// b.Run(fmt.Sprintf("%v", count), func(b *testing.B) { singleNumberB(f, b) })
	// b.Run(fmt.Sprintf("%v", count), func(b *testing.B) { singleNumberB(f, b) })
	// b.Run(fmt.Sprintf("%v", count), func(b *testing.B) { singleNumberB(f, b) })
	// b.Run(fmt.Sprintf("%v", count), func(b *testing.B) { singleNumberB(f, b) })
}

func TestSingleNumberD(t *testing.T) {
	f := singleNumberMapDiv
	t.Run(fmt.Sprintf("%v", count), func(t *testing.T) { singleNumberT(f, t) })
	t.Run(fmt.Sprintf("%v", count), func(t *testing.T) { singleNumberT(f, t) })
	t.Run(fmt.Sprintf("%v", count), func(t *testing.T) { singleNumberT(f, t) })
	t.Run(fmt.Sprintf("%v", count), func(t *testing.T) { singleNumberT(f, t) })
	t.Run(fmt.Sprintf("%v", count), func(t *testing.T) { singleNumberT(f, t) })
	t.Run(fmt.Sprintf("%v", count), func(t *testing.T) { singleNumberT(f, t) })
	t.Run(fmt.Sprintf("%v", count), func(t *testing.T) { singleNumberT(f, t) })
}

func TestMain(m *testing.M) {
	// activate benchmarking if required
	testing.Init()

	// parse command line flags for test
	flag.Parse()

	// run and exit
	os.Exit(m.Run())
}
