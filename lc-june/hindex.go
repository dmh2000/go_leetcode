package lcjune

import "fmt"

func hmin(i int, j int) int {
	if i > j {
		return j
	}
	return i
}

func hIndex1(citations []int) int {
	var j int
	var h int
	a := []int{}
	j = 1
	for i := len(citations); i > 0; i-- {
		h = hmin(citations[i-1], j)
		a = append(a, h)
		fmt.Printf("(%v,%v,%v),", citations[i-1], j, h)
		j++
	}
	fmt.Println(a)

	h = a[0]
	for i := len(citations) - 1; i > 0; i-- {
		if a[i-1] < a[i] {
			h = a[i]
			break
		}
	}

	return h
}

func hIndex(citations []int) int {
	return 0
}
