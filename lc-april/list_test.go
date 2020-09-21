package lc

import (
	"testing"
)

func push(list *MyIntList, val int) {
	list.PushFront(val)

}

func pop(list *MyIntList, t *testing.T) int {
	val, ok := list.PopFront()
	if !ok {
		t.Error("PopFront failed")
	}
	return val
}

func popval(list *MyIntList, val int, t *testing.T) int {
	v, ok := list.PopFront()
	if !ok {
		t.Error("PopFront failed")
	}
	if v != val {
		t.Errorf("PopFront value %v should be %v", v, val)
	}
	return val
}

func TestList1(t *testing.T) {
	var list *MyIntList
	list = NewMyIntList()

	if list.Head() != nil {
		t.Error("Head should be nil")
	}

	if list.Tail() != nil {
		t.Error("Tail should be nil")
	}

	// push 100 values in increasing value
	for i := 0; i < 100; i++ {
		push(list, i)
	}

	// pop from head, comes back in reverse value
	for i := 99; i >= 1; i-- {
		v := list.Head().data.(int)
		if v != i {
			t.Errorf("List head has wrong value, is %v should be %v", v, i)
		}
		popval(list, i, t)
	}

	var v int

	// one item remaining, should be 0
	v = list.Head().data.(int)
	if v != 0 {
		t.Error("Head should be 0 but is ", v)
	}

	// one item remaining, should be 0
	v = list.Tail().data.(int)
	if v != 0 {
		t.Error("Tail should be 0 but is ", v)
	}

	// one item remaining, shoudld be 0
	v = popval(list, 0, t)

	// nothing left
	if list.Head() != nil {
		t.Error("Head should be nil")
	}

	// nothing left
	if list.Tail() != nil {
		t.Error("Tail should be nil")
	}

}
