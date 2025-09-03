package main

import "fmt"

func main() {

	mySlice := make([]int, 100)
	for i := range len(mySlice) {
		mySlice[i] = i
	}

	fmt.Println(exists(mySlice, 50))

}

func exists(arr []int, num int) bool {
	firstP := 0
	secondP := len(arr) - 1
	for firstP <= secondP {
		midIdx := (firstP + secondP) / 2
		currentN := arr[midIdx]
		if currentN == num {
			return true
		} else if currentN > num {
			secondP = midIdx - 1
		} else {
			firstP = midIdx + 1
		}
	}

	return false
}
