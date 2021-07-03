package main

import "fmt"

/*
	The trie data structure is a structure based on a tree and is commonly used to find words
	its time complexity is based on what we are looking for, i.e, looking for the word "complex"
	will take at least O(m) where m is the length of the word "complex"
*/

const (
	// alphabetSize is the number of possible character in the trie
	alphabetSize = 26
)

// Node represents each node in the  trie
type Node struct {
	children [alphabetSize]*Node
	isEnd    bool
}

// Trie represents a trie and has a pointer to the root node
type Trie struct {
	root *Node
}

// InitTrie will create new Trie
func InitTrie() *Trie {
	return &Trie{root: &Node{}}
}

// Insert will take in a word and add it to the trie
func (t *Trie) Insert(w string) {
	currentNode := t.root
	for _, letter := range w {
		charIdx := letter - 'a'
		if currentNode.children[charIdx] == nil {
			currentNode.children[charIdx] = &Node{}
		}

		currentNode = currentNode.children[charIdx]
	}
	currentNode.isEnd = true

}

// Search will take in a word and RETURN true if that word is included in the trie
func (t *Trie) Search(w string) bool {
	currentNode := t.root
	for _, letter := range w {
		charIdx := letter - 'a'
		if currentNode.children[charIdx] == nil {
			return false
		}

		currentNode = currentNode.children[charIdx]
	}
	return currentNode.isEnd
}

func main() {

	trie := InitTrie()
	trie.Insert("aragorn")
	trie.Insert("besides")
	fmt.Println(trie.root)
	fmt.Println(trie.Search("aragorn"))
}
