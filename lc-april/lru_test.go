package lc

import "testing"

func TestLru1(t *testing.T) {
	lru := Constructor(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	v := lru.Get(1)
	if v != 1 {
		t.Error(v, " should be ", 1)
	}
	lru.Put(3, 3)
	v = lru.Get(2)
	if v != -1 {
		t.Error(v, " should be ", -1)
	}
	lru.Put(4, 4)
	v = lru.Get(1)
	if v != -1 {
		t.Error(v, " should be ", -1)
	}
	v = lru.Get(3)
	if v != 3 {
		t.Error(v, " should be ", 3)
	}
	v = lru.Get(4)
	if v != 4 {
		t.Error(v, " should be ", 4)
	}

}

func TestLru2(t *testing.T) {
	lru := Constructor(1)
	lru.Put(1, 1)
	lru.Put(2, 2)
	v := lru.Get(1)
	if v != -1 {
		t.Error(v, " should be ", -1)
	}
	v = lru.Get(2)
	if v != 2 {
		t.Error(v, " should be ", 2)
	}
	lru.Put(3, 3)
	v = lru.Get(3)
	if v != 3 {
		t.Error(v, " should be ", 3)
	}

}

func TestLru3(t *testing.T) {
	lru := Constructor(10)

	for i := 0; i < 11; i++ {
		lru.Put(i, i+100)
	}
	v := lru.Get(1)
	if v != 101 {
		t.Error(101, " should be ", 101)
	}

	lru.Put(12, 12)
}

func TestLru4(t *testing.T) {
	lru := Constructor(2)
	lru.Put(2, 1)
	lru.Put(2, 2)
	v := lru.Get(2)
	if v != 2 {
		t.Error(v, " should be ", 2)
	}
	lru.Put(1, 1)
	lru.Put(4, 1)
	v = lru.Get(2)
	if v != -1 {
		t.Error(v, " should be ", -1)
	}
}

func TestLru5(t *testing.T) {
	lru := Constructor(2)
	lru.Put(2, 1)
	lru.Put(1, 1)
	lru.Put(2, 3)
	lru.Put(4, 1)
	v := lru.Get(1)
	if v != -1 {
		t.Error(v, " should be ", -1)
	}
	v = lru.Get(2)
	if v != 3 {
		t.Error(v, " should be ", 3)
	}
}

func TestLru6(t *testing.T) {
	// [[2],[2],[2,6],[1],[1,5],[1,2],[1],[2]]
	lru := Constructor(2)
	v := lru.Get(2)
	if v != -1 {
		t.Error(v, " should be ", -1)
	}
	lru.Put(2, 6)
	v = lru.Get(1)
	if v != -1 {
		t.Error(v, " should be ", -1)
	}
	lru.Put(1, 5)
	lru.Put(1, 2)
	v = lru.Get(1)
	if v != 2 {
		t.Error(v, " should be ", 2)
	}
	v = lru.Get(2)
	if v != 6 {
		t.Error(v, " should be ", 2)
	}
}
