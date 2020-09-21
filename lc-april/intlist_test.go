package lc

import "testing"

func TestIntList1(t *testing.T) {
	list := NewIntList()

	list.PushFront(1)
	count := list.Len()
	if count != 1 {
		t.Error(count, " should be ", 1)
	}
	list.PushFront(2)
	count = list.Len()
	if count != 2 {
		t.Error(count, " should be ", 2)
	}
	list.PushFront(3)
	count = list.Len()
	if count != 3 {
		t.Error(count, " should be ", 3)
	}
	list.PrintFwd()
	list.PrintBak()
}

func TestIntList2(t *testing.T) {
	list := NewIntList()

	list.PushBack(1)
	count := list.Len()
	if count != 1 {
		t.Error(count, " should be ", 1)
	}
	list.PushBack(2)
	count = list.Len()
	if count != 2 {
		t.Error(count, " should be ", 2)
	}
	list.PushBack(3)
	count = list.Len()
	if count != 3 {
		t.Error(count, " should be ", 3)
	}
	list.PrintFwd()
	list.PrintBak()
}

func TestIntList3(t *testing.T) {
	list := NewIntList()

	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)
	list.PrintBak()

	v, ok := list.PopBack()
	if !ok {
		t.Error(ok, " should be ", true)
	}
	if v != 3 {
		t.Error(v, " should be ", 3)
	}

	v, ok = list.PopBack()
	if !ok {
		t.Error(ok, " should be ", true)
	}
	if v != 2 {
		t.Error(v, " should be ", 2)
	}

	v, ok = list.PopBack()
	if !ok {
		t.Error(ok, " should be ", true)
	}
	if v != 1 {
		t.Error(v, " should be ", 1)
	}

	v, ok = list.PopBack()
	if ok {
		t.Error(ok, " should be ", false)
	}
	if v != 0 {
		t.Error(v, " should be ", 0)
	}

	if !list.Empty() {
		t.Error("list should be empty")
	}

	if list.Len() != 0 {
		t.Error("list length should be 0")
	}
}

func TestIntList4(t *testing.T) {
	list := NewIntList()

	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)
	list.PrintBak()
	list.PrintFwd()
	v, ok := list.PopFront()
	if !ok {
		t.Error(ok, " should be ", true)
	}
	if v != 1 {
		t.Error(v, " should be ", 1)
	}

	v, ok = list.PopFront()
	if !ok {
		t.Error(ok, " should be ", true)
	}
	if v != 2 {
		t.Error(v, " should be ", 2)
	}

	v, ok = list.PopFront()
	if !ok {
		t.Error(ok, " should be ", true)
	}
	if v != 3 {
		t.Error(v, " should be ", 3)
	}

	v, ok = list.PopBack()
	if ok {
		t.Error(ok, " should be ", false)
	}
	if v != 0 {
		t.Error(v, " should be ", 0)
	}

	if !list.Empty() {
		t.Error("list should be empty")
	}

	if list.Len() != 0 {
		t.Error("list length should be 0")
	}
}
func TestIntList5(t *testing.T) {
	var node *IntListNode
	list := NewIntList()

	// node in between
	list.PushFront(1)
	node = list.PushFront(2)
	list.PushFront(3)
	list.PrintFwd()
	list.Remove(node)
	list.PrintFwd()
	list.PrintBak()
}

func TestIntList6(t *testing.T) {
	var node *IntListNode
	list := NewIntList()

	// node at tail
	node = list.PushFront(1)
	list.PushFront(2)
	list.PushFront(3)
	list.PrintFwd()
	list.Remove(node)
	count = list.Len()
	if count != 2 {
		t.Error(count, " should be ", 2)
	}
	list.PrintFwd()
	list.PrintBak()
}

func TestIntList7(t *testing.T) {
	var node *IntListNode
	list := NewIntList()

	// node at head
	list.PushFront(1)
	list.PushFront(2)
	node = list.PushFront(3)
	list.PrintFwd()
	list.Remove(node)
	count = list.Len()
	if count != 2 {
		t.Error(count, " should be ", 2)
	}
	list.PrintFwd()
	list.PrintBak()
}
