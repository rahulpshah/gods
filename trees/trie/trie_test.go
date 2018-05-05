package trie

import (
	"testing"
)

func buildTrie(trie *Trie, words ...string) *Trie {
	for _, word := range words {
		trie.Add(word)
	}
	return trie
}

func TestTrie(t *testing.T) {
	trie := buildTrie(NewTrie(), "rahul", "ra")
	cases := []struct {
		searchWord string
		present    bool
	}{
		{"rahul", true},
		{"rahu", false},
		{"ra", true},
		{"prashant", false},
		{"shah", false},
		{"", false},
	}
	for _, c := range cases {
		isPresent := trie.Has(c.searchWord)
		if c.present != isPresent {
			t.Errorf("Expected %v, Got %v", c.present, isPresent)
		}
	}
	if trie.Size() != 2 {
		t.Errorf("Expected 2, Got %v", trie.Size())
	}
	if trie.Empty() {
		t.Errorf("Expected false, Got true")
	}
	trie.Clear()
	if !trie.Empty() {
		t.Errorf("Expected true, Got false")
	}
	trie.Clear()
	if trie.Size() != 0 {
		t.Errorf("Expected 0, Got %v", trie.Size())
	}
	trie.Add("shah")
	if !trie.Has("shah") {
		t.Errorf("Expected true, Got false")
	}
}
