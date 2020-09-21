package lc

import (
	"testing"
)

func TestMyList1(t *testing.T) {
	var i = 0
	mlist := NewMyList()

	mlist.Init()
	for i = 0; i < 10; i++ {
		mlist.PushBack(i)
	}

	i = 0
	for e := mlist.Front(); e != nil; e = e.Next() {
		if e.Value != i {
			t.Error(e.Value, " should be ", i)
		}
		i++
	}

	i = 9
	for e := mlist.Back(); e != nil; e = e.Prev() {
		if e.Value != i {
			t.Error(e.Value, " should be ", i)
		}
		i--
	}

}

func TestMyInsertAfter(t *testing.T) {
	var a []int
	var i int
	var e *MyElement
	mlist := NewMyList()

	mlist.Init()

	a = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < 10; i++ {
		mlist.PushBack(a[i])
	}

	if mlist.Len() != len(a) {
		t.Error(mlist.Len(), " should be ", len(a))
	}

	e = mlist.InsertAfter(10, mlist.Front())
	if e.Value != 10 {
		t.Error(e.Value, " should be ", 10)
	}

	i = 0
	a = []int{0, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for e := mlist.Front(); e != nil; e = e.Next() {
		if e.Value != a[i] {
			t.Error(e.Value, " should be ", a[i])
		}
		i++
	}

	if mlist.Len() != len(a) {
		t.Error(mlist.Len(), " should be ", len(a))
	}

	e = mlist.InsertAfter(11, mlist.Back())
	if e.Value != 11 {
		t.Error(e.Value, " should be ", 10)
	}

	i = 0
	a = []int{0, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 11}
	for e := mlist.Front(); e != nil; e = e.Next() {
		if e.Value != a[i] {
			t.Error(e.Value, " should be ", a[i])
		}
		i++
	}
	if mlist.Len() != len(a) {
		t.Error(mlist.Len(), " should be ", len(a))
	}

}

func TestMyInsertBefore(t *testing.T) {
	var a []int
	var i int
	var e *MyElement
	mlist := NewMyList()

	mlist.Init()

	a = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < 10; i++ {
		mlist.PushBack(a[i])
	}

	e = mlist.InsertBefore(10, mlist.Front())
	if e.Value != 10 {
		t.Error(e.Value, " should be ", 10)
	}

	i = 0
	a = []int{10, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for e := mlist.Front(); e != nil; e = e.Next() {
		if e.Value != a[i] {
			t.Error(e.Value, " should be ", a[i])
		}
		i++
	}

	e = mlist.InsertBefore(11, mlist.Back())
	if e.Value != 11 {
		t.Error(e.Value, " should be ", 11)
	}

	i = 0
	a = []int{10, 0, 1, 2, 3, 4, 5, 6, 7, 8, 11, 9}
	for e := mlist.Front(); e != nil; e = e.Next() {
		if e.Value != a[i] {
			t.Error(e.Value, " should be ", a[i])
		}
		i++
	}
}

func TestRemove(t *testing.T) {
	var a []int
	var e *MyElement
	mlist := NewMyList()

	mlist.Init()

	a = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < 10; i++ {
		mlist.PushBack(a[i])
	}

	// remove '1'
	e = mlist.Front().Next()
	if e.Value != 1 {
		t.Error(e.Value, " should be ", 1)
	}

	mlist.Remove(e)

	if mlist.Len() != 9 {
		t.Error(mlist.Len(), " should be ", 9)
	}

	// check list with '1' removed
	a = []int{0, 2, 3, 4, 5, 6, 7, 8, 9}
	e = mlist.Front()
	for i := 0; i < len(a); i++ {
		if e.Value != a[i] {
			t.Error(e.Value, " should be ", a[i])
		}
		e = e.Next()
	}

}

func TestMoveAfter(t *testing.T) {
	var a []int
	var e *MyElement
	var f *MyElement
	mlist := NewMyList()

	a = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < 10; i++ {
		mlist.PushBack(a[i])
	}

	// move 0 after 1
	e = mlist.Front()
	f = mlist.Front().Next()
	mlist.MoveAfter(e, f)

	// test it
	a = []int{1, 0, 2, 3, 4, 5, 6, 7, 8, 9}
	e = mlist.Front()
	for i := 0; i < len(a); i++ {
		if e.Value != a[i] {
			t.Error(e.Value, " should be ", a[i])
		}
		e = e.Next()
	}
}

func TestMoveBefore(t *testing.T) {
	var a []int
	var e *MyElement
	var f *MyElement
	mlist := NewMyList()

	a = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < 10; i++ {
		mlist.PushBack(a[i])
	}

	// move 0 after 1
	e = mlist.Front()
	f = mlist.Front().Next()
	mlist.MoveBefore(f, e)

	// test it
	a = []int{1, 0, 2, 3, 4, 5, 6, 7, 8, 9}
	e = mlist.Front()
	for i := 0; i < len(a); i++ {
		if e.Value != a[i] {
			t.Error(e.Value, " should be ", a[i])
		}
		e = e.Next()
	}
}

func TestMoveFront(t *testing.T) {
	var a []int
	var e *MyElement
	mlist := NewMyList()

	a = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < 10; i++ {
		mlist.PushBack(a[i])
	}

	// move 0 after 1
	e = mlist.Back()
	mlist.MoveToFront(e)

	// test it
	a = []int{9, 0, 1, 2, 3, 4, 5, 6, 7, 8}
	e = mlist.Front()
	for i := 0; i < len(a); i++ {
		if e.Value != a[i] {
			t.Error(e.Value, " should be ", a[i])
		}
		e = e.Next()
	}
}

func TestMoveBack(t *testing.T) {
	var a []int
	var e *MyElement
	mlist := NewMyList()

	a = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < len(a); i++ {
		mlist.PushBack(a[i])
	}

	// move 0 after 1
	e = mlist.Front()
	mlist.MoveToBack(e)

	// test it
	a = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	e = mlist.Front()
	for i := 0; i < len(a); i++ {
		if e.Value != a[i] {
			t.Error(e.Value, " should be ", a[i])
		}
		e = e.Next()
	}
}

func TestMyPushFrontList(t *testing.T) {
	var a []int
	var b []int
	var e *MyElement
	m1 := NewMyList()
	m2 := NewMyList()

	// populate m1
	m1.Init()
	a = []int{0, 1, 2, 3, 4}
	for i := 0; i < len(a); i++ {
		m1.PushBack(a[i])
	}

	m2.Init()
	b = []int{5, 6, 7, 8, 9}
	for i := 0; i < len(b); i++ {
		m2.PushBack(b[i])
	}

	// append m2 to m1
	m2.PushFrontList(m1)

	if m2.Len() != len(a)+len(b) {
		t.Error(m1.Len(), " should be ", len(a)+len(b))
	}

	//test it
	a = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	e = m2.Front()
	for i := 0; i < len(a); i++ {
		if e.Value != a[i] {
			t.Error(e.Value, " should be ", a[i])
		}
		e = e.Next()
	}
}

func TestMyPushBackList(t *testing.T) {
	var a []int
	var b []int
	var e *MyElement
	m1 := NewMyList()
	m2 := NewMyList()

	// populate m1
	m1.Init()
	a = []int{0, 1, 2, 3, 4}
	for i := 0; i < len(a); i++ {
		m1.PushBack(a[i])
	}

	m2.Init()
	b = []int{5, 6, 7, 8, 9}
	for i := 0; i < len(b); i++ {
		m2.PushBack(b[i])
	}

	// append m2 to m1
	m1.PushBackList(m2)

	if m1.Len() != len(a)+len(b) {
		t.Error(m1.Len(), " should be ", len(a)+len(b))
	}
	//test it
	a = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	e = m1.Front()
	for i := 0; i < len(a); i++ {
		if e.Value != a[i] {
			t.Error(e.Value, " should be ", a[i])
		}
		e = e.Next()
	}
}
