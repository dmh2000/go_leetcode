package lcmay

// type int2a [][]int

// func (a int2a) Len() int           { return len(a) }
// func (a int2a) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
// func (a int2a) Less(i, j int) bool { return compare(a[i], a[j]) < 0 }

// func mergeIntersect2(A [][]int, B [][]int) [][]int {
// 	var r [][]int

// 	r = make([][]int, 0, len(A)+len(B))
// 	r = append(r, A...)
// 	r = append(r, B...)

// 	sort.Stable(int2a(r))
// 	return r
// }

// // =====================================================

// type int2 []int

// func (a int2) compare(b []int) int {
// 	i := a[0] - b[0]
// 	j := a[1] - b[1]

// 	r := 0
// 	if i < 0 {
// 		r = i
// 	} else if i == 0 {
// 		r = j
// 	} else {
// 		r = i
// 	}
// 	return r
// }

// func mergeIntersect1(A [][]int, B [][]int) [][]int {
// 	var r [][]int

// 	r = make([][]int, 0, len(A)+len(B))
// 	// merge the two arrays in sorted order
// 	i := 0
// 	j := 0
// 	for i < len(A) && j < len(B) {
// 		a := int2(A[i])
// 		b := int2(B[j])
// 		k := a.compare(b)
// 		if k == 0 {
// 			// keep one and discard the other
// 			r = append(r, a)
// 			i++
// 			j++
// 		} else if k < 0 {
// 			// a is less
// 			r = append(r, a)
// 			i++
// 		} else {
// 			r = append(r, b)
// 			j++
// 		}
// 	}
// 	// append rest of A if any
// 	for i < len(A) {
// 		r = append(r, A[i])
// 		i++
// 	}
// 	for j < len(B) {
// 		r = append(r, B[j])
// 		j++
// 	}

// 	return r
// }
// ==============================================

// func compare(a []int, b []int) int {
// 	i := a[0] - b[0]
// 	j := a[1] - b[1]

// 	r := 0
// 	if i < 0 {
// 		r = i
// 	} else if i == 0 {
// 		r = j
// 	} else {
// 		r = i
// 	}
// 	return r
// }

// func mergeIntersect(A [][]int, B [][]int) [][]int {
// 	var r [][]int

// 	r = make([][]int, 0, len(A)+len(B))
// 	// merge the two arrays in sorted order
// 	i := 0
// 	j := 0
// 	for i < len(A) && j < len(B) {
// 		a := A[i]
// 		b := B[j]
// 		k := compare(a, b)
// 		if k == 0 {
// 			// keep one and discard the other
// 			r = append(r, a)
// 			i++
// 			j++
// 		} else if k < 0 {
// 			// a is less
// 			r = append(r, a)
// 			i++
// 		} else {
// 			r = append(r, b)
// 			j++
// 		}
// 	}
// 	// append rest of A if any
// 	for i < len(A) {
// 		r = append(r, A[i])
// 		i++
// 	}
// 	for j < len(B) {
// 		r = append(r, B[j])
// 		j++
// 	}

// 	return r
// }

// func intervalIntersectionX(A [][]int, B [][]int) [][]int {
// 	var a []int
// 	var b []int
// 	var c []int

// 	var r [][]int
// 	var s [][]int
// 	// merge the two lists
// 	r = mergeIntersect(A, B)
// 	if len(r) <= 1 {
// 		return r
// 	}

// 	s = make([][]int, 0, len(A)+len(B))
// 	fmt.Println(r)
// 	i := 0
// 	for {
// 		// take next two items from the list
// 		if i >= len(r) {
// 			break
// 		}
// 		a = r[i]

// 		if i+1 >= len(r) {
// 			break
// 		}
// 		b = r[i+1]

// 		// find the largest [0] and smallest [1]
// 		c = []int{0, 0}
// 		// is there an overlap

// 		// overlap
// 		c[0] = b[0]
// 		c[1] = a[1]

// 		fmt.Println(a, b, c)

// 		if c[0] <= c[1] {
// 			s = append(s, c)
// 		}

// 		i++
// 	}

// 	return s
// }

func intervalIntersection(A [][]int, B [][]int) [][]int {
	var a int
	var b int
	var lo int
	var hi int
	var s [][]int

	s = make([][]int, 0, len(A)+len(B))

	a = 0
	b = 0
	for a < len(A) && b < len(B) {
		// max of the [0] element
		if A[a][0] > B[b][0] {
			lo = A[a][0]
		} else {
			lo = B[b][0]
		}

		// min of the [1] element
		if A[a][1] < B[b][1] {
			hi = A[a][1]
		} else {
			hi = B[b][1]
		}

		// if there is a positive overlap
		if lo <= hi {
			// keep it
			s = append(s, []int{lo, hi})
		}

		// skip the smaller of the endpoints
		if A[a][1] < B[b][1] {
			a++
		} else {
			b++
		}
	}

	return s
}
