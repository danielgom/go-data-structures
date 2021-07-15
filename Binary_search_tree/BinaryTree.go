package main

import "fmt"

var count int

// Both non-recursive and recursive creations of BST have a O(log(n)) time creation, however
// the recursive methods have O(log(n)) space complexity because of the stack trace having this on memory

// Node represents the components of a binary search tree
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// Insert will add a node to the tree
// the key to add should not be already in the tree
func (n *Node) Insert(k int) {

	this := n

	for {
		if k > this.Key {
			if this.Right == nil {
				this.Right = &Node{Key: k}
				break
			} else {
				this = this.Right
			}
		} else if k < this.Key {
			if this.Left == nil {
				this.Left = &Node{Key: k}
				break
			} else {
				this = this.Left
			}
		}
	}
}

// Search will take in a key value
// and return true if there is a node with that value
func (n *Node) Search(k int) bool {

	this := n

	for this != nil {
		if k > this.Key {
			this = this.Right
		} else if k < this.Key {
			this = this.Left
		} else {
			return true
		}
	}
	return false
}

func (n *Node) Remove(k int) {
	n.removeHelper(k, nil)
}

func (n *Node) removeHelper(k int, parent *Node) {

}

func main() {

	tree := &Node{
		Key:   100,
		Left:  nil,
		Right: nil,
	}

	tree.Insert(101)
	tree.Insert(52)
	tree.Insert(19)
	tree.Insert(76)
	tree.Insert(150)
	tree.Insert(310)
	fmt.Println(tree.Search(52))
	fmt.Println(count)

}
