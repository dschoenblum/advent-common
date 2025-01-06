package common

import "iter"

type trieNode[T any] struct {
	children map[byte]*trieNode[T]
	values   []T
}

func newTrieNode[T any]() *trieNode[T] {
	return &trieNode[T]{
		children: make(map[byte]*trieNode[T]),
	}
}

type Trie[T any] struct {
	root *trieNode[T]
}

func NewTrie[T any]() *Trie[T] {
	return &Trie[T]{
		root: newTrieNode[T](),
	}
}

func (t *Trie[T]) Add(entry string, value T) {
	node := t.root
	for i := 0; i < len(entry); i++ {
		char := entry[i]
		if node.children[char] == nil {
			node.children[char] = newTrieNode[T]()
		}
		node = node.children[char]
	}
	node.values = append(node.values, value)
}

func (t *Trie[T]) AllPrefixes(input string) iter.Seq2[string, []T] {
	return func(yield func(string, []T) bool) {
		node := t.root
		for i := 0; i < len(input); i++ {
			char := input[i]
			if node.children[char] == nil {
				break
			}
			node = node.children[char]
			if len(node.values) > 0 {
				if !yield(input[:i+1], node.values) {
					return
				}
			}
		}
	}
}
