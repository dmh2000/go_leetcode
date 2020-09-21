package lcmay

import "testing"

func TestTrie1(t *testing.T) {
	var tr Trie
	var b bool
	tr = TrieConstructor()

	tr.Insert("apple")

	b = tr.Search("apple")
	if !b {
		t.Error(b, " should be ", true)
	}

	b = tr.StartsWith("app")
	if !b {
		t.Error(b, " should be ", true)
	}

	b = tr.Search("app")
	if b {
		t.Error(b, " should be ", false)
	}

	tr.Insert("app")

	b = tr.StartsWith("app")
	if !b {
		t.Error(b, " should be ", true)
	}

	b = tr.Search("app")
	if !b {
		t.Error(b, " should be ", true)
	}

	b = tr.Search("appreciate")
	if b {
		t.Error(b, " should be ", false)
	}

	b = tr.StartsWith("appr")
	if b {
		t.Error(b, " should be ", false)
	}

	tr.Insert("appreciate")

	b = tr.Search("appreciate")
	if !b {
		t.Error(b, " should be ", true)
	}

	b = tr.StartsWith("appr")
	if !b {
		t.Error(b, " should be ", true)
	}

	tr.Print()
}

func TestTrie2(t *testing.T) {
	var tr Trie
	var b bool
	var words []string

	words = []string{
		"apple",
		"appreciate",
		"beta",
		"clorox",
		"cloak",
		"zero",
		"zeroize",
	}

	tr = TrieConstructor()

	for _, v := range words {
		tr.Insert(v)
	}

	for _, v := range words {
		b = tr.Search(v)
		if !b {
			t.Error(b, " should be ", true, " for ", v)
		}
	}

}
