package search

import "fmt"

/*
	implement trie tree by slice
*/

type sliceTrieNode struct {
	pass, end int
	nexts     []*sliceTrieNode
}

func newTrieNode1(size int) *sliceTrieNode {
	return &sliceTrieNode{
		pass:  0,
		end:   0,
		nexts: make([]*sliceTrieNode, size),
	}
}

type SliceTrieTree struct {
	startChar rune
	pathSize  int
	root      *sliceTrieNode
}

func NewSliceTrieTree(startChar rune, size int) *SliceTrieTree {
	return &SliceTrieTree{
		startChar: startChar,
		pathSize:  size,
		root:      newTrieNode1(size),
	}
}

func (tt *SliceTrieTree) indexOfChar(r rune) int {
	return int(r - tt.startChar)
}

func (tt *SliceTrieTree) charOfIndex(index int) rune {
	return rune(int(tt.startChar) + index)
}

func (tt *SliceTrieTree) validString(word string) bool {
	for _, r := range word {
		if tt.indexOfChar(r) >= tt.pathSize {
			return false
		}
	}
	return true
}

func (tt *SliceTrieTree) Insert(word string) {
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

func (tt *SliceTrieTree) Delete(word string) {
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

func (tt *SliceTrieTree) Search(word string) int {
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

func (tt *SliceTrieTree) PrefixCount(pre string) int {
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

/*
	implement trie tree by map
*/

type mapTrieNode struct {
	pass, end int
	nexts     map[rune]*mapTrieNode
}

func newMapTrieNode() *mapTrieNode {
	return &mapTrieNode{
		pass:  0,
		end:   0,
		nexts: make(map[rune]*mapTrieNode, 0),
	}
}

type MapTrieTree struct {
	root *mapTrieNode
}

func NewMapTrieTree() *MapTrieTree {
	return &MapTrieTree{
		root: newMapTrieNode(),
	}
}

func (t *MapTrieTree) Insert(word string) {
	if len(word) < 1 {
		return
	}
	node := t.root
	node.pass++
	for _, c := range word {
		if _, ok := node.nexts[c]; !ok {
			node.nexts[c] = newMapTrieNode()
		}
		node = node.nexts[c]
		node.pass++
	}
	node.end++
}

func (t *MapTrieTree) Search(word string) bool {
	if len(word) < 1 {
		return false
	}
	node := t.root
	for _, c := range word {
		if _, ok := node.nexts[c]; !ok {
			return false
		}
		node = node.nexts[c]
	}
	return node.end > 0
}

func (t *MapTrieTree) Delete(word string) {
	if t.Search(word) {
		node := t.root
		node.pass--
		for _, c := range word {
			node = node.nexts[c]
			node.pass--
		}
		node.end--
	}
}

func (t *MapTrieTree) PrefixCount(word string) int {
	if t.Search(word) {
		node := t.root
		for _, c := range word {
			if _, ok := node.nexts[c]; !ok {
				return 0
			}
			node = node.nexts[c]
		}
		return node.pass
	}
	return 0
}
