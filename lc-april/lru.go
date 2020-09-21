package lc

import (
	"container/list"
	"fmt"
)

// LruItem ...
type LruItem struct {
	key   int
	value int
	node  *list.Element
}

// LRUCache ...
type LRUCache struct {
	hash   map[int]LruItem // store values in hash
	keyq   *list.List      // store keys in the queue
	cap    int
	length int
}

// Constructor ...
func Constructor(capacity int) LRUCache {
	lru := LRUCache{}
	lru.cap = capacity
	lru.hash = make(map[int]LruItem)
	lru.keyq = list.New()
	lru.length = 0
	return lru
}

func (lru *LRUCache) moveToFront(key int) {
	item, ok := lru.hash[key]
	if !ok {
		panic("Invalid key")
	}
	lru.keyq.MoveToFront(item.node)
}

// Get ...
func (lru *LRUCache) Get(key int) int {
	// the hash contains the value
	item, ok := lru.hash[key]
	if !ok {
		return -1
	}

	// move the KEY to the front of the queue
	lru.keyq.MoveToFront(item.node)

	lru.print("get")
	return item.value
}

// Put ...
func (lru *LRUCache) Put(key int, value int) {
	item, ok := lru.hash[key]

	// if already in cache
	if ok {
		// update the value
		item.value = value
	} else {
		// new item
		item = LruItem{value: value, key: key, node: nil}
		// if at capacity, remove the least recently used value
		if lru.length == lru.cap {
			// get the last queue element
			node := lru.keyq.Back()
			// save it values
			k := node.Value.(int)
			// remove it from the queue
			lru.keyq.Remove(node)
			// remove it from the hash table
			delete(lru.hash, k)
		} else {
			lru.length++
		}
	}

	// if it was already in the cache
	if ok {
		// move it to front of the queue
		lru.keyq.MoveToFront(item.node)
	} else {
		// push it on the front
		lru.keyq.PushFront(item.key)
	}

	// update its node field
	item.node = lru.keyq.Front()

	// update the hash map
	lru.hash[item.key] = item

	lru.print("put")
}

func (lru *LRUCache) print(s string) {
	fmt.Printf("%v : ", s)
	p := lru.keyq.Front()
	for p != nil {
		fmt.Printf("%v,", p.Value)
		p = p.Next()
	}
	fmt.Println(lru.hash)
}
