type Trie struct {
	children map[byte]*Trie
	isWord   bool
}

func Constructor() Trie {
	return Trie{
		children: make(map[byte]*Trie),
		isWord:   false,
	}
}

func (this *Trie) Insert(word string) {
	if word == "" {
		this.isWord = true
		return
	}
	node, ok := this.children[word[0]]
	if !ok {
		node = &Trie{
			children: make(map[byte]*Trie),
			isWord:   false,
		}
		this.children[word[0]] = node
	}
	node.Insert(word[1:])
}

func (this *Trie) Search(word string) bool {
	if word == "" {
		return this.isWord
	}

	node, ok := this.children[word[0]]
	if !ok {
		return false
	}

	return node.Search(word[1:])
}

func (this *Trie) StartsWith(prefix string) bool {
	if prefix == "" {
		return true
	}

	node, ok := this.children[prefix[0]]
	if !ok {
		return false
	}

	return node.StartsWith(prefix[1:])
}