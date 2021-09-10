package main

import (
	"fmt"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

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

// Remove will remove and balance the BinarySearchTree
func (n *Node) Remove(k int) {
	n.removeHelper(k, nil)
}

// removeHelper is the central method in order to remove a node from a BST
func (n *Node) removeHelper(key int, parent *Node) {

	currentNode := n

 	for currentNode != nil {
 		fmt.Println(currentNode.Key)
 		if currentNode.Key > key {
 			parent = currentNode
 			currentNode = currentNode.Left
		} else if currentNode.Key < key {
			parent = currentNode
			currentNode = currentNode.Right
		} else {
			if currentNode.Left != nil && currentNode.Right != nil {
				currentNode.Key = currentNode.Right.getMinVal()
				currentNode.Right.removeHelper(currentNode.Key, currentNode)
			} else if parent == nil {
				if currentNode.Left != nil {
					currentNode.Key = currentNode.Left.Key
					currentNode.Left = currentNode.Left.Left
				} else if currentNode.Right != nil {
					currentNode.Key = currentNode.Right.Key
					currentNode.Right = currentNode.Right.Right
				} else {
					currentNode = nil
				}
			} else if parent.Left == currentNode {
				if currentNode.Left != nil {
					parent.Left = currentNode.Left
				} else {
					parent.Left = currentNode.Right
				}
			} else if parent.Right == currentNode {
				if currentNode.Left != nil {
					parent.Right = currentNode.Left
				} else {
					parent.Right = currentNode.Right
				}
			}
			break
		}
	}
}

// getMinVal returns the minimum value/key from the current node, the minimum value of a BST
// is the leftmost value in the last level (height 0)
func (n *Node) getMinVal() int {
	current := n

	for current != nil {
		current = current.Left
	}

	return current.Key
}

// PrintInOrder prints into the console the ordered BST
func (n *Node) PrintInOrder() {

	current := n

	if current.Left != nil {
		current.Left.PrintInOrder()
	}

	fmt.Print(" ", current.Key)

	if current.Right != nil {
		current.Right.PrintInOrder()
	}

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
	tree.Remove(150)
	fmt.Println(tree.Search(52))
	fmt.Println(count)

	fmt.Println("printing ordered bst")
	tree.PrintInOrder()
	fmt.Println()

	intervals := [][]int{{10, 20}, {21, 32},{1, 5},{2, 10},{13, 15},{6, 9},{8, 17}}
	fmt.Println(intervals)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	fmt.Println(intervals)

	var longInterval []int
	longInterval = append(longInterval, 1,2,3)
	fmt.Println(longInterval)
	fmt.Println(longInterval[0:0])

	var notes []string

	notes = []string{"10.0% higher than in-store",
	"5.0% lower than in-store",
	"Same as in-store"}
	fmt.Println(notes)

	fmt.Println(strings.Contains(notes[2], "Same"))
	compile := regexp.MustCompile("[^0-9.]")
	fmt.Println(compile.MatchString("10.0% higher than in-store"))
	fmt.Println(compile.ReplaceAllString("10.0% higher than in-store", ""))

	allString := compile.ReplaceAllString("10.133% higher than in-store", "")
	float, err := strconv.ParseFloat(allString, 10)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(float)

}
