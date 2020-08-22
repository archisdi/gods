package tree

import "fmt"

// Node represent an element in tree
type Node struct {
	Data  *int
	Left  *Node
	Right *Node
}

// Insert a new node into tree
func (n *Node) Insert(value int) {
	if n.Data == nil {
		n.Data = &value
	} else if *n.Data >= value {
		if n.Left == nil {
			n.Left = &Node{
				Data: &value,
			}
		} else {
			n.Left.Insert(value)
		}
	} else {
		if n.Right == nil {
			n.Right = &Node{
				Data: &value,
			}
		} else {
			n.Right.Insert(value)
		}
	}
}

// TraverseInorder prints every node in tree from left, top, to right
func (n *Node) TraverseInorder() {
	if n.Left != nil {
		n.Left.TraverseInorder()
	}
	fmt.Println(*n.Data)
	if n.Right != nil {
		n.Right.TraverseInorder()
	}
}

// TraversePreorder prints every node in tree from top, left, to right
func (n *Node) TraversePreorder() {
	fmt.Println(*n.Data)
	if n.Left != nil {
		n.Left.TraversePreorder()
	}
	if n.Right != nil {
		n.Right.TraversePreorder()
	}
}

// TraversePostorder prints every node in tree from left, right, to top
func (n *Node) TraversePostorder() {
	if n.Left != nil {
		n.Left.TraversePostorder()
	}
	if n.Right != nil {
		n.Right.TraversePostorder()
	}
	fmt.Println(*n.Data)
}

// Lowest returns min value node in tree
func (n *Node) Lowest() int {
	if n.Left != nil {
		return n.Left.Lowest()
	}
	return *n.Data
}

// Highest returns max value in tree
func (n *Node) Highest() int {
	if n.Right != nil {
		return n.Right.Highest()
	}
	return *n.Data
}

// Exist search a value within tree
func (n *Node) Exist(value int) bool {
	if *n.Data == value {
		return true
	}
	if *n.Data >= value {
		if n.Left != nil {
			return n.Left.Exist(value)
		}
		return false
	}
	if n.Right != nil {
		return n.Right.Exist(value)
	}
	return false
}

// New returns new tree instance
func New(elements []int) Node {
	tree := Node{}
	for _, element := range elements {
		tree.Insert(element)
	}
	return tree
}
