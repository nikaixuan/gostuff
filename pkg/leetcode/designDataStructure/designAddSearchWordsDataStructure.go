package designDataStructure

// Need trieNode
type WordDictionary struct {
	head *trieNode
}

type trieNode struct {
	child []*trieNode
	end   bool
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	dict := WordDictionary{}
	dict.head = NewtrieNode()
	return dict
}

func NewtrieNode() *trieNode {
	n := trieNode{}
	n.child = make([]*trieNode, 26)
	n.end = false
	return &n
}

func (this *WordDictionary) AddWord(word string) {
	n := this.head
	for _, w := range word {
		if n.child[w-'a'] == nil {
			n.child[w-'a'] = NewtrieNode()
		}
		n = n.child[w-'a']
	}
	n.end = true
}

func (this *WordDictionary) Search(word string) bool {
	return this.head.search(word)
}

func (node *trieNode) search(word string) bool {
	if len(word) == 0 {
		return node.end
	}
	if word[0] == '.' {
		for _, n := range node.child {
			if n != nil && n.search(word[1:]) {
				return true
			}
		}
	} else {
		return node.child[word[0]-'a'] != nil && node.child[word[0]-'a'].search(word[1:])
	}
	return false
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
