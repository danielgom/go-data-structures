package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
	list *linkedList
}

type linkedList struct {
	// head is the first element of the linkedList
	head *Node
	len  int
}

func (l *linkedList) prepend(n *Node) {
	second := l.head
	l.head = n
	l.head.list = l
	l.head.next = second
	l.len++
}

func (l *linkedList) insert(n *Node, index int) {
	if l.len == 0 {
		return
	}

	if index == 1 {
		l.prepend(n)
		return
	}

	previousNode := l.head
	currentPosition := 2
	for currentPosition < index {
		previousNode = previousNode.next
		currentPosition++
	}

	n.next = previousNode.next
	n.list = l
	previousNode.next = n
}

func (l *linkedList) deleteWithValue(v int) {
	if l.len == 0 {
		return
	}

	if l.head.data == v {
		l.head = l.head.next
		l.len--
		return
	}

	previous := l.head
	for previous.next.data != v {
		if previous.next.next == nil {
			return
		}
		previous = previous.next
	}

	previous.next = previous.next.next
	l.len--
}

func (l *linkedList) printListData() {
	for e := l.head; e != nil; e = e.next {
		fmt.Print(e.data, " ")
	}
	fmt.Println()
}

func main() {
	l := &linkedList{}
	node1 := &Node{data: 48}
	node2 := &Node{data: 18}
	node3 := &Node{data: 16}
	node4 := &Node{data: 11}
	l.prepend(node1)
	l.prepend(node2)
	l.prepend(node3)
	l.prepend(node4)
	l.printListData()

	l.deleteWithValue(11)
	l.insert(&Node{data: 100}, 4)
	l.printListData()
}
