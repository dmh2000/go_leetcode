package lc

import "container/list"

// MyType ...

// MyList ...
type MyList struct {
	// underlying container/list
	mylist *list.List
}

// MyElement ...
type MyElement struct {
	// list.List Element pointer
	element *list.Element

	// needs this so client doesn't need to do e.element.Value
	// if you change this type, then change the instances where
	// e.Value and the 'int' parameters are used
	Value int
}

// Next ...tested
func (m *MyElement) Next() *MyElement {
	e := m.element.Next()
	if e == nil {
		return nil
	}
	return &MyElement{element: e, Value: e.Value.(int)}
}

// Prev ...tested
func (m *MyElement) Prev() *MyElement {
	e := m.element.Prev()
	if e == nil {
		return nil
	}
	return &MyElement{element: e, Value: e.Value.(int)}
}

// NewMyList ...tested
func NewMyList() *MyList {
	return &MyList{mylist: list.New()}
}

// Back ...tested
func (m *MyList) Back() *MyElement {
	e := m.mylist.Back()
	return &MyElement{element: e, Value: e.Value.(int)}
}

// Front ...tested
func (m *MyList) Front() *MyElement {
	e := m.mylist.Front()
	return &MyElement{element: e, Value: e.Value.(int)}
}

// Init ...tested
func (m *MyList) Init() *MyList {
	m.mylist.Init()
	return m
}

// InsertAfter ...tested
func (m *MyList) InsertAfter(v int, mark *MyElement) *MyElement {
	e := m.mylist.InsertAfter(v, mark.element)
	return &MyElement{element: e, Value: e.Value.(int)}
}

// InsertBefore ...tested
func (m *MyList) InsertBefore(v int, mark *MyElement) *MyElement {
	e := m.mylist.InsertBefore(v, mark.element)
	return &MyElement{element: e, Value: e.Value.(int)}
}

// Len ...tested
func (m *MyList) Len() int {
	return m.mylist.Len()
}

// MoveAfter ...tested
func (m *MyList) MoveAfter(e, mark *MyElement) {
	m.mylist.MoveAfter(e.element, mark.element)
}

// MoveBefore ...tested
func (m *MyList) MoveBefore(e, mark *MyElement) {
	m.mylist.MoveBefore(e.element, mark.element)
}

// MoveToBack ...tested
func (m *MyList) MoveToBack(e *MyElement) {
	m.mylist.MoveToBack(e.element)
}

// MoveToFront ...testd
func (m *MyList) MoveToFront(e *MyElement) {
	m.mylist.MoveToFront(e.element)
}

// PushBack ...tested
func (m *MyList) PushBack(v int) *MyElement {
	e := m.mylist.PushBack(v)
	return &MyElement{element: e}
}

// PushBackList ...
func (m *MyList) PushBackList(other *MyList) {
	m.mylist.PushBackList(other.mylist)
}

// PushFront ...tested
func (m *MyList) PushFront(v int) *MyElement {
	e := m.mylist.PushFront(v)
	return &MyElement{element: e}
}

// PushFrontList ...tested
func (m *MyList) PushFrontList(other *MyList) {
	m.mylist.PushFrontList(other.mylist)
}

// Remove ...tested
func (m *MyList) Remove(e *MyElement) int {
	xe := m.mylist.Remove(e.element)
	return xe.(int)
}
