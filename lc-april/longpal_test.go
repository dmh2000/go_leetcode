package lc

import (
	"fmt"
	"testing"
)

func pals() []string {
	return []string{
		"babab",
		"babbab",
	}
}

func notPals() []string {
	return []string{
		"bac",
		"ba",
	}
}

func words() []string {
	return []string{
		"abb",
		"babaddtattarrattatddetartrateedredividerb",
		"a",
		"baba",
		"babac",
		"babbab",
		"baabdaadbaab",
		"baabdaadbaac",
	}
}
func TestIsPalindrome(t *testing.T) {
	var pal string
	var w []string

	w = pals()
	for _, s := range w {
		if !isPalindrome(s) {
			t.Error("should be a palindrome palindrome", s, pal)
		}
	}
	w = notPals()
	for _, s := range w {
		if isPalindrome(s) {
			t.Error("should not be a palindrome", s, pal)
		}
	}
}

func TestLongestPalindrome(t *testing.T) {
	var w []string

	p := ""
	for i := 0; i < 100000; i++ {
		w = words()
		for _, s := range w {
			p = longestPalindrome(s)
			// fmt.Println(s, p)
		}
	}
	fmt.Println(p)
}
