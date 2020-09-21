package lc

// FirstUnique ...
type FirstUnique struct {
	// queue []int
	count map[int]int
	value int
	queue []int
}

// Constructor

// Unique ...
func Unique(nums []int) FirstUnique {
	f := FirstUnique{}
	f.count = make(map[int]int)
	f.queue = make([]int, 0)
	f.value = -1

	for _, v := range nums {
		f.Add(v)
	}
	return f
}

// ShowFirstUnique ...
func (f *FirstUnique) ShowFirstUnique() int {
	return f.value
}

// Add ...
func (f *FirstUnique) Add(value int) {
	// add it to count
	f.count[value]++

	if f.value == -1 {
		// no previous unique and this one is unique
		// add it
		if f.count[value] == 1 {
			f.value = value
			f.queue = make([]int, 0)
			f.queue = append(f.queue, []int{value}...)
		} else {
			// truncate it
			f.queue = make([]int, 0)
		}
	} else {
		// add new value to the queue
		f.queue = append(f.queue, []int{value}...)

		// find the current unique value
		if f.count[f.value] > 1 {
			f.value = -1
			// find next unique
			for i, v := range f.queue {
				if f.count[v] == 1 {
					f.value = v
					f.queue = f.queue[i:]
					break
				}
			}
		}
	}
	// fmt.Println(value, f)
}

/**
 * Your FirstUnique object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.ShowFirstUnique();
 * obj.Add(value);
 */
