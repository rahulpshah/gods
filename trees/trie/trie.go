package trie

// Trie is the representation of the ADT
type Trie struct {
	root   *Node
	size   int
	values []interface{}
}

// Size returns the number of words in the trie
func (t *Trie) Size() int {
	return len(t.values)
}

// Values returns the all the words in the trie
func (t *Trie) Values() []interface{} {
	return t.values
}

// Clear removes all the nodes
func (t *Trie) Clear() {
	t.values = nil
	t.root = NewNode()
}

// Empty checks if the trie is empty
func (t *Trie) Empty() bool {
	return t.Size() == 0
}

// Node represents a single node in the trie
type Node struct {
	isEnd bool
	mp    map[rune]*Node
}

// NewTrie is constructor for trie
func NewTrie() *Trie {
	trie := &Trie{}
	trie.root = NewNode()
	return trie
}

// NewNode is constructor for node
func NewNode() *Node {
	node := &Node{}
	node.mp = make(map[rune]*Node, 0)
	return node
}

// Add adds a word to the trie
func (t *Trie) Add(word string) {
	p := t.root
	for _, c := range word {
		if _, ok := p.mp[c]; !ok {
			p.mp[c] = NewNode()
		}
		p = p.mp[c]
	}
	p.isEnd = true
	t.values = append(t.values, word)

}

// Has checks if the word is in the trie
func (t *Trie) Has(word string) bool {
	p := t.root
	for _, c := range word {
		if _, ok := p.mp[c]; !ok {
			return false
		}
		p = p.mp[c]
	}
	return p.isEnd
}
