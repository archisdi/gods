package main

import (
	"ds/search"
)

func main() {
	// tree := tree.New([]int{10, 15, 5, 22, 8, 1, 4})
	// graph := graph.New([]int{1, 2, 3, 4, 5, 6, 7, 8}, []graph.Edge{{1, 4}, {1, 5}, {2, 5}, {2, 6}, {3, 1}, {3, 2}, {3, 6}, {6, 7}, {7, 8}, {5, 6}})
	// linkedList := linkedlist.New([]int{1, 9, 2, 8, 3, 7, 5, 6})
	// linkedStack := list.NewLinkedStack([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	// heap := tree.NewMinHeap([]int{10, 15, 20, 17, 25})

	list := search.List{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 16, 20, 50, 56, 67, 78, 89, 90, 100, 245}
	list.BinarySearch(245)
}
