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
		if k >= this.Key {
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

// InOrderTraverse will print the BST in order
func (n *Node) InOrderTraverse() {

	var pre *Node
	this := n

	for this != nil {
		if this.Left == nil {
			fmt.Println(this.Key)
			this = this.Right
		} else {
			pre = this.Left
			for pre.Right != nil && pre.Right != this {
				pre = pre.Right
			}
			if pre.Right == nil {
				pre.Right = this
				this = this.Left
			} else {
				pre.Right = nil
				fmt.Println(this.Key)
				this = this.Right
			}
		}
	}
}

func (n *Node) Remove(k int) {
	n.removeHelper(k, nil)
}

func (n *Node) removeHelper(k int, parent *Node) {

}

func main() {

	tree := &Node{
		Key:   10,
		Left:  nil,
		Right: nil,
	}

	tree.Insert(5)
	tree.Insert(15)
	tree.Insert(2)
	tree.Insert(5)
	tree.Insert(1)
	tree.Insert(22)
	fmt.Println(tree.Search(52))

	tree.InOrderTraverse()

	intervals := [][]int{{10, 20}, {21, 32}, {1, 5}, {2, 10}, {13, 15}, {6, 9}, {8, 17}}
	fmt.Println(intervals)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	fmt.Println(intervals)

	var longInterval []int
	longInterval = append(longInterval, 1, 2, 3)
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
