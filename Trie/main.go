package main

import "fmt"

type TrieNode struct {
	childern map[rune]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

func (t *Trie) Insert(word string) {
	node := t.root
	for _, str := range word {
		if _, exists := t.childern[str]; !exists {
			node.childern[char] = &TrieNode{childern: make(map[rune]*TrieNode)}
		}
		node = node.childern[char]
	}
	node.isEnd = true
}

func (t *Trie) Search(word string) bool {
	node := t.root
	for _, str := range word {
		if _, exists := node.childern[str]; !exists {
			return false
		}
		node = node.childern[str]
	}
	return node.isEnd
}

func (t *Trie) StartsWith(word string)bool {
	node := t.root
	for _, str := range word {
		if _, exists := node.childern[str]; !exists {
			return false
		}
		node = node.childern[str]
	}

	return true
}

func main() {
	trie := &Trie{}
	trie.Insert("Apple")
	trie.Insert("Banana")
	fmt.Println(trie.Search("Apple"))

}
