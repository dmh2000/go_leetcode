package lc

import (
	"fmt"
	"testing"
)

func anagramCases() [][]string {
	a := make([][]string, 0)
	a = append(a, []string{"eat", "tea", "tan", "ate", "nat", "bat"})

	return a
}
func TestAnagram1(t *testing.T) {
	cases := anagramCases()
	for _, v := range cases {
		a := groupAnagrams(v)
		fmt.Println(a)
	}
}
