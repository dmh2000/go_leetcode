package lcjune

import (
	"math/rand"
)

type pickSolution struct {
	table []int
	scale float64
	wsum  float64
	size  int
}

func pickConstructor(w []int) pickSolution {
	var s pickSolution
	s.size = len(w)
	s.scale = 10000.0

	// sum the weights
	s.wsum = 0
	for _, v := range w {
		s.wsum += float64(v)
	}

	// make a table of 10^5
	s.table = make([]int, int(s.scale))

	// distribute the weights
	dt := (1.0 / s.scale)
	idx := 0
	for i := 0; i < s.size; i++ {
		weight := (float64(w[i]) / s.wsum)
		for weight > 0.0 && idx < len(s.table) {
			s.table[idx] = i
			idx++
			weight -= dt
			if weight < 0.0001 {
				weight = 0
			}
		}
	}
	return s
}

func (p *pickSolution) pickIndex() int {
	var f float64
	var r int
	var i int

	// get a random number in the range [0.0,1.0)
	f = rand.Float64()

	// convert it to an scaled integer [0..scale]
	f = f * p.scale

	// cast to int
	r = int(f)

	// return the selected index
	i = p.table[r]

	return i
}
