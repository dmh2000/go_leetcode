package lcjune

import (
	"fmt"
	"testing"
)

func TestRandomSet1(t *testing.T) {
	var r RandomizedSet
	var b bool
	var v int

	r = RConstructor()

	b = r.Insert(1)
	if !b {
		t.Error(b, "should be", true)
	}

	b = r.Remove(2)
	if b {
		t.Error(b, "should be", true)
	}

	b = r.Insert(2)
	if !b {
		t.Error(b, "should be", true)
	}

	v = r.GetRandom()
	if !(v == 1 || v == 2) {
		t.Error(b, "should be 1 or 2")
	}

	b = r.Remove(1)
	if !b {
		t.Error(b, "should be", true)
	}

	b = r.Insert(2)
	if b {
		t.Error(b, "should be", false)
	}

	v = r.GetRandom()
	if v != 2 {
		t.Error(b, "should be", 2)
	}

}

func TestRandomSet2(t *testing.T) {
	var r RandomizedSet
	var b bool
	var v int

	r = RConstructor()

	b = r.Remove(0)
	if b {
		t.Error(b, "should be", false)
	}

	b = r.Remove(0)
	if b {
		t.Error(b, "should be", false)
	}

	b = r.Insert(0)
	if !b {
		t.Error(b, "should be", true)
	}

	v = r.GetRandom()
	if v != 0 {
		t.Error(v, "should be ", 0)
	}

	b = r.Remove(0)
	if !b {
		t.Error(b, "should be", true)
	}

	b = r.Insert(0)
	if !b {
		t.Error(b, "should be", true)
	}
}

func TestRandomSet3(t *testing.T) {
	var r RandomizedSet
	var b bool
	var v int

	r = RConstructor()

	b = r.Insert(0)
	if !b {
		t.Error(b, "should be", true)
	}

	b = r.Insert(1)
	if !b {
		t.Error(b, "should be", true)
	}

	b = r.Remove(0)
	if !b {
		t.Error(b, "should be", true)
	}

	b = r.Insert(2)
	if !b {
		t.Error(b, "should be", true)
	}

	b = r.Remove(1)
	if !b {
		t.Error(b, "should be", true)
	}

	v = r.GetRandom()
	if v != 2 {
		t.Error(v, "should be ", 2)
	}
}

func TestRandomSet4(t *testing.T) {
	var r RandomizedSet
	var b bool
	var v int

	r = RConstructor()

	b = r.Insert(3)
	if !b {
		t.Error(b, "should be", true)
	}

	b = r.Remove(3)
	if !b {
		t.Error(b, "should be", true)
	}

	b = r.Remove(0)
	if b {
		t.Error(b, "should be", false)
	}

	b = r.Insert(3)
	if !b {
		t.Error(b, "should be", true)
	}

	v = r.GetRandom()
	if v != 3 {
		t.Error(v, "should be ", 3)
	}

	b = r.Remove(0)
	if b {
		t.Error(b, "should be", false)
	}
}

var ops []string = []string{"RandomizedSet", "insert", "insert", "remove", "insert",
	"insert", "insert", "remove", "remove", "insert", "remove",
	"insert", "insert", "insert", "insert", "insert", "getRandom",
	"insert", "remove", "insert", "insert"}

var arg [][]int = [][]int{{}, {3}, {-2}, {2}, {1}, {-3}, {-2}, {-2}, {3}, {-1}, {-3}, {1},
	{-2}, {-2}, {-2}, {1}, {}, {-2}, {0}, {-3}, {1}}

var ans = []bool{true, true, true, false, true, true, false, true, true, true, true, false, true, false, false, false, true, false, false, true, false}

func TestRandomSet5(t *testing.T) {
	var r RandomizedSet
	var b bool
	var v int
	for i := 0; i < len(ops); i++ {
		switch ops[i] {
		case "RandomizedSet":
			r = RConstructor()
			b = true
		case "insert":
			v = arg[i][0]
			b = r.Insert(v)
			if b != ans[i] {
				fmt.Printf("%v:%v:%v,", i, b, v)
				t.Error(b, "should be", ans[i])
			}
		case "remove":
			v = arg[i][0]
			b = r.Remove(v)
			if b != ans[i] {
				fmt.Printf("%v:%v:%v,", i, b, v)
				t.Error(b, "should be", ans[i])
			}
		case "getRandom":
			v = r.GetRandom()
			fmt.Printf("%v,", v)
		}
	}
	fmt.Println()
}
