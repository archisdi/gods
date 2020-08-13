package main

import (
	"ds/tree"
	"fmt"
)

func main() {
	tree := tree.New([]int{10, 15, 5, 22, 8, 1, 4})

	// fmt.Println(tree.Lowest())
	// fmt.Println(tree.Highest())
	fmt.Println(tree.Exist(5))
}
