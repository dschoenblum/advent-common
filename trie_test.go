package common

import (
	"slices"
	"testing"
)

func TestTrie(t *testing.T) {
	trie := NewTrie[bool]()
	trie.Add("a", true)
	trie.Add("ab", true)
	trie.Add("abcd", true)
	trie.Add("xyz", true)

	var prefixes []string
	for prefix, _ := range trie.AllPrefixes("abcdef") {
		prefixes = append(prefixes, prefix)
	}

	if !slices.Equal(prefixes, []string{"a", "ab", "abcd"}) {
		t.Errorf("Expected a, ab, abcd, got %v", prefixes)
	}
}
