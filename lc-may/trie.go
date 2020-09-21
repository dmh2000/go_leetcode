package lcmay

import "fmt"

type trieNode struct {
	val  [26]byte      // records the character, indicates which letters are valid
	term [26]bool      // indicates a word is terminated at this level
	next [26]*trieNode // pointers to next level
}

// Trie ...
type Trie struct {
	head *trieNode
}

// TrieConstructor ...
// Initialize your data structure here.
func TrieConstructor() Trie {
	t := Trie{}
	t.head = &trieNode{}

	return t
}

// recursive insertion
func (t *Trie) insert(node *trieNode, word string) {
	// quit when no more characters
	if len(word) == 0 {
		return
	}

	// get the next character and offset to lower case
	index := word[0] - 'a'

	// check if already populated
	if node.val[index] == 0 {
		// not yet populated,
		node.val[index] = word[0]
		node.next[index] = &trieNode{}
	}

	// if this is the last character
	if len(word) == 1 {
		// mark it as a word
		node.term[index] = true
	}

	// recurse into rest of word
	t.insert(node.next[index], word[1:])
}

// Insert a word into the trie.
func (t *Trie) Insert(word string) {
	// quit on empty string
	if len(word) == 0 {
		return
	}
	// insert recursively
	t.insert(t.head, word)
}

// recursive search
func (t *Trie) search(node *trieNode, word string) bool {
	index := word[0] - 'a'

	// base case
	if len(word) == 1 {
		return node.term[index]
	}

	if node.next[index] == nil {
		return false
	}

	return t.search(node.next[index], word[1:])
}

// Search Returns if the word is in the trie
func (t *Trie) Search(word string) bool {
	// quit on empty string
	if len(word) == 0 {
		return false
	}

	// search recursively
	return t.search(t.head, word)
}

// recursive startsWith
func (t *Trie) startsWith(node *trieNode, prefix string) bool {
	index := prefix[0] - 'a'

	// stop on last character
	if len(prefix) == 1 {
		if node.val[index] != 0 {
			return true
		}
	}

	// don't recurse further
	if node.next[index] == nil {
		return false
	}

	// keep searching
	return t.startsWith(node.next[index], prefix[1:])
}

// StartsWith Returns if there is any word in the trie that starts with the given prefix.
func (t *Trie) StartsWith(prefix string) bool {
	if len(prefix) == 0 {
		return false
	}

	return t.startsWith(t.head, prefix)
}
func (t *Trie) print(node *trieNode) {
	if node == nil {
		fmt.Println()
	}

	for i := 0; i < 26; i++ {
		if node.val[i] != 0 {
			fmt.Printf("%v", string(node.val[i]))
			t.print(node.next[i])
		}
	}
}

// Print ...
func (t *Trie) Print() {
	t.print(t.head)
	fmt.Println()
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
