package lc

import (
	"fmt"
	"testing"
)

func TestMinStack1(t *testing.T) {
	mstack := NewMinStack()

	for i := 0; i < 10; i++ {
		mstack.Push(i)
	}

	for i := 9; i >= 0; i-- {
		top := mstack.Top()
		if top != i {
			t.Error(top, " should be ", i)
		}
		min := mstack.GetMin()
		if min != 0 {
			t.Error(min, " should be ", 0)
		}
		mstack.Pop()

		length := mstack.Len()
		if length != i {
			t.Error(length, " should be ", i)
		}
	}
}

func TestMinStack2(t *testing.T) {
	mstack := NewMinStack()

	j := 9
	for i := 0; i < 10; i++ {
		mstack.Push(j)
		j--
	}

	j = 9
	for i := 0; i < 10; i++ {
		top := mstack.Top()
		if top != i {
			t.Error(top, " should be ", i)
		}
		min := mstack.GetMin()
		if min != i {
			t.Error(min, " should be ", i)
		}
		mstack.Pop()

		length := mstack.Len()
		if length != j {
			t.Error(length, " should be ", j)
		}
		j--
	}
}

func TestMinStack3(t *testing.T) {
	// ["MinStack","push","push","push","getMin","pop","top","getMin"]
	// [[],[-2],[0],[-3],[],[],[],[]]
	mstack := NewMinStack()
	mstack.Push(-2)
	mstack.Push(0)
	mstack.Push(-3)
	min := mstack.GetMin()
	if min != -3 {
		t.Error(min, " should be ", -3)
	}
	mstack.Pop()
	x := mstack.Top()
	if x != 0 {
		t.Error(x, " should be ", 0)
	}
	min = mstack.GetMin()
	if min != -2 {
		t.Error(min, " should be ", -2)
	}
}

func remove(mstack *MinStack, count int, b *testing.B) {
	length := mstack.Len()
	for i := 0; i < length; i++ {
		mstack.Pop()
	}
}

func populateUp(count int, b *testing.B) {
	mstack := NewMinStack()
	for i := 0; i < count; i++ {
		mstack.Push(i)
	}
	remove(&mstack, count, b)
}

func populateDown(count int, b *testing.B) {
	mstack := NewMinStack()
	for i := count - 1; i >= 0; i-- {
		mstack.Push(i)
	}
	remove(&mstack, count, b)
}

func BenchmarkMinstack(b *testing.B) {
	var count int
	count = 1024
	for i := 0; i < 6; i++ {
		// pstack = populateUp(count)
		b.Run(fmt.Sprintf("Down : %v", count), func(b *testing.B) { populateDown(count, b) })
		b.Run(fmt.Sprintf("UP : %v", count), func(b *testing.B) { populateUp(count, b) })
		count *= 2
	}
}
