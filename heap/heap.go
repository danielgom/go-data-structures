package main

import (
	"fmt"
)

/*
	Heaps can be expressed as a complete tree, all the levels are full except for the lowest level,
	where some nodes can be empty but you have all the nods at the left.
	Max heap means that the largest number stays in the root position, Min heap is just the opposite

	The time complexity of a heap can be expressed by O(h) where h means the high of the heap, which
	is basically the size of the array , we can replace O(h) to O(log n) because the height and the number of
	indices have a logarithmic relation
*/

// MaxHeap struct holds a slice that holds the array

type MaxHeap struct {
	array []int
}

// Insert add and element to the heap
func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

// Extract returns the largest key and removes it from the heap, if the slice is empty
// -1 is returned
func (h *MaxHeap) Extract() {

	if len(h.array) == 0 {
		fmt.Println("cannot extract element because slice length is 0 or nil")
	}

	l := len(h.array) - 1

	// Place last element to the root
	h.array[0] = h.array[l]

	// Reduce slice size
	h.array = h.array[:l]

	h.maxHeapifyDown(0)

}

// maxHeapifyUp will heapify from bottom to top
func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

// maxHeapifyDown will heapify from top to bottom
func (h *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	// loop while index has at least one child
	for l <= lastIndex {
		// when left is the only child
		if l == lastIndex {
			childToCompare = l
		} else if h.array[l] > h.array[r] { // when left child is larger than the right child
			childToCompare = l
		} else { // when right child is larger
			childToCompare = r
		}
		// compare array value of current index to larger child and swap if smaller parent
		if h.array[index] < h.array[childToCompare] {
			// swapping if the child is larger than the parent
			h.swap(index, childToCompare)
			// making the new index the child swapped
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			// if the child is not larger, it means that it has found its place
			return
		}
	}

}

// swap keys in the array
func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

// get the parent index
func parent(i int) int {
	return (i - 1) / 2
}

//get the left child index
func left(p int) int {
	return p*2 + 1
}

// get the right child index
func right(p int) int {
	return p*2 + 2
}

func main() {

	m := &MaxHeap{}
	fmt.Println(m)
	buildHeap := []int{10, 20, 30, 5, 7, 9, 11, 13, 15, 17}
	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Println(m)
	}

	for i := 0; i < 5; i++ {
		m.Extract()
		fmt.Println(m)
	}
}
