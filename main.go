package main

import (
	"ds/linkedlist"
)

func main() {
	// tree := tree.New([]int{10, 15, 5, 22, 8, 1, 4})
	// graph := graph.New([]int{1, 2, 3, 4, 5, 6, 7, 8}, []graph.Edge{{1, 4}, {1, 5}, {2, 5}, {2, 6}, {3, 1}, {3, 2}, {3, 6}, {6, 7}, {7, 8}, {5, 6}})

	list := linkedlist.New([]int{1, 9, 2, 8, 3, 7, 5, 6})

	list.Append(10)
	list.Delete(1)
	list.Prepend(69)
	list.Delete(10)
	list.Print()
}
