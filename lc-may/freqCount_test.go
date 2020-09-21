package lcmay

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
)

func TestFreqSort(t *testing.T) {
	var s string
	var r string
	s = frequencySortS("tree")
	r = frequencySortB("tree")
	if s != "eetr" {
		t.Error(s, "sould be", "eetr")
	}
	fmt.Println(s, r)

	s = frequencySortB("cccaaa")
	if s != "aaaccc" {
		t.Error(s, "sould be", "cccaaa")
	}
	fmt.Println(s)

	s = frequencySortB("Aabb")
	if s != "bbAa" {
		t.Error(s, "sould be", "bbAa")
	}
	fmt.Println(s)
}

var bstring string

func BenchmarkFreqSortS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		frequencySortS(bstring)
	}
}

func BenchmarkFreqSortB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		frequencySortB(bstring)
	}
}

func TestMain(m *testing.M) {
	for i := 0; i < 100000; i++ {
		bstring += string(byte(rand.Int31n(256)))
	}
	os.Exit(m.Run())
}

/*
go test -test.bench=. -test.run=xxxxx


Given a string, sort it in decreasing order based on the frequency of characters.

Example 1:

Input:
"tree"

Output:
"eert"

Explanation:
'e' appears twice while 'r' and 't' both appear once.
So 'e' must appear before both 'r' and 't'. Therefore "eetr" is also a valid answer.
Example 2:

Input:
"cccaaa"

Output:
"cccaaa"

Explanation:
Both 'c' and 'a' appear three times, so "aaaccc" is also a valid answer.
Note that "cacaca" is incorrect, as the same characters must be together.
Example 3:

Input:
"Aabb"

Output:
"bbAa"

Explanation:
"bbaA" is also a valid answer, but "Aabb" is incorrect.
Note that 'A' and 'a' are treated as two different characters.
*/
