package lcjune

import "math/rand"

// RandomizedSet ...
type RandomizedSet struct {
	m map[int]int
	a []int
}

// RConstructor : Initialize your data structure here
func RConstructor() RandomizedSet {
	r := RandomizedSet{}
	r.m = make(map[int]int)
	r.a = make([]int, 0, 100000)
	return r
}

// Insert : Inserts a value to the set. Returns true if the set did not already contain the specified element
func (r *RandomizedSet) Insert(val int) bool {
	// is it already there?
	_, ok := r.m[val]
	if ok {
		// already in there
		return false
	}

	// no, add it

	// index where it will be inserted
	index := len(r.a)
	// add to map
	r.m[val] = index
	// add to array
	r.a = append(r.a, val)

	// return false since it was not already in the map
	return true
}

// Remove : Removes a value from the set. Returns true if the set contained the specified element.
func (r *RandomizedSet) Remove(val int) bool {

	// is the list empty?
	if len(r.a) == 0 {
		return false
	}

	// is it there?
	index, ok := r.m[val]
	if !ok {
		// no, just return false
		return false
	}

	// remove from list

	// is there only one in the list
	if len(r.a) == 1 {
		// just start over
		r.a = make([]int, 0, 100000)
		r.m = make(map[int]int)
		return true
	}

	// copy last element to index
	v := r.a[len(r.a)-1]
	r.a[index] = v

	// update its position in the map
	r.m[v] = index

	// remove last element
	r.a = r.a[:len(r.a)-1]

	// remove from map
	delete(r.m, val)

	return true
}

// GetRandom : Get a random element from the set
func (r *RandomizedSet) GetRandom() int {

	// Intn doesn't work if length is 0
	if len(r.a) == 1 {
		return r.a[0]
	}

	// get random index in range [0,n) and return the value
	index := rand.Intn(len(r.a))
	return r.a[index]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
