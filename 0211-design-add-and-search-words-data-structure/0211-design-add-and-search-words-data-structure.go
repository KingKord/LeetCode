type WordDictionary struct {
	children map[byte]*WordDictionary
	isWord   bool
}
func Constructor() WordDictionary {
	return WordDictionary{
		children: make(map[byte]*WordDictionary),
		isWord:   false,
	}
}

func (this *WordDictionary) AddWord(word string) {
	if word == "" {
		this.isWord = true
		return
	}
	node, ok := this.children[word[0]]
	if !ok {
		node = &WordDictionary{
			children: make(map[byte]*WordDictionary),
			isWord:   false,
		}
		this.children[word[0]] = node
	}
	
	node.AddWord(word[1:])
}

func (this *WordDictionary) Search(word string) bool {
	if word == "" {
		return this.isWord
	}
	
	if word[0] == '.' {
		for _, dictionary := range this.children {
			if dictionary.Search(word[1:]) {
				return true
			}
		}
		return false
	}
	node, ok := this.children[word[0]]
	if !ok {
		return false
	}

	return node.Search(word[1:])
}