package lcjune

import (
	"fmt"
	"testing"
)

func TestRpick1(t *testing.T) {

	var s pickSolution
	var p int
	var a []int

	s = pickConstructor([]int{1})

	for i := 0; i < 10; i++ {
		p = s.pickIndex()
		if p != 0 {
			t.Error(p, "should be", 0)
		}
	}

	s = pickConstructor([]int{1, 3})

	a = make([]int, 2)
	for i := 0; i < 100; i++ {
		p = s.pickIndex()
		a[p]++
	}
	fmt.Println(a)

}

func TestRpick2(t *testing.T) {

	var s pickSolution
	var p int
	var a []int

	s = pickConstructor([]int{20000, 10000, 10000, 10000, 10000, 10000})

	a = make([]int, s.size)
	for i := 0; i < 10000; i++ {
		p = s.pickIndex()
		a[p]++
	}
	fmt.Println(a)
}

func TestRpick3(t *testing.T) {

	var s pickSolution
	var p int
	var a []int

	a = make([]int, 0)
	for i := 0; i < 100000; i++ {
		a = append(a, 100000)
	}
	s = pickConstructor(a)

	a = make([]int, s.size)
	for i := 0; i < 10000; i++ {
		p = s.pickIndex()
		a[p]++
	}
}
