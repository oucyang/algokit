package search

import "fmt"

type trieNode1 struct {
	pass, end int
	nexts     []*trieNode1
}

func newTrieNode1(size int) *trieNode1 {
	return &trieNode1{
		pass:  0,
		end:   0,
		nexts: make([]*trieNode1, size),
	}
}

type TrieTree1 struct {
	startChar rune
	pathSize  int
	root      *trieNode1
}

func NewTrieTree1(startChar rune, size int) *TrieTree1 {
	return &TrieTree1{
		startChar: startChar,
		pathSize:  size,
		root:      newTrieNode1(size),
	}
}

func (tt *TrieTree1) indexOfChar(r rune) int {
	return int(r - tt.startChar)
}

func (tt *TrieTree1) charOfIndex(index int) rune {
	return rune(int(tt.startChar) + index)
}

func (tt *TrieTree1) validString(word string) bool {
	for _, r := range word {
		if tt.indexOfChar(r) >= tt.pathSize {
			return false
		}
	}
	return true
}

func (tt *TrieTree1) Insert(word string) {
	if len(word) < 1 {
		return
	}
	if tt.validString(word) == false {
		panic(fmt.Sprintf("'%s' contains char out of range.", word))
	}
	node := tt.root
	node.pass++
	for _, r := range word {
		index := tt.indexOfChar(r)
		if node.nexts[index] == nil {
			node.nexts[index] = newTrieNode1(tt.pathSize)
		}
		node = node.nexts[index]
		node.pass++
	}
	node.end++
}

func (tt *TrieTree1) Delete(word string) {
	if tt.Search(word) > 0 {
		node := tt.root
		node.pass--
		for _, r := range word {
			index := tt.indexOfChar(r)
			node.nexts[index].pass--
			if node.nexts[index].pass == 0 {
				node.nexts[index] = nil
				return
			}
			node = node.nexts[index]
		}
		node.end--
	}
}

func (tt *TrieTree1) Search(word string) int {
	if len(word) < 1 {
		return 0
	}
	node := tt.root
	for _, r := range word {
		index := tt.indexOfChar(r)
		if node.nexts[index] == nil {
			return 0
		}
		node = node.nexts[index]
	}
	return node.end
}

func (tt *TrieTree1) PrefixCount(pre string) int {
	if len(pre) < 1 {
		return 0
	}
	node := tt.root
	for _, r := range pre {
		index := tt.indexOfChar(r)
		if node.nexts[index] == nil {
			return 0
		}
		node = node.nexts[index]
	}
	return node.pass
}

type trieNode2 struct {
	pass, end int
	nexts     map[rune]*trieNode2
}
