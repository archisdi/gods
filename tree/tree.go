package tree

import "fmt"

// Node ...
type Node struct {
	Data  *int
	Left  *Node
	Right *Node
}

// IsLeaf ...
func (n *Node) IsLeaf() bool {
	return n.Left == nil && n.Right == nil
}

// Insert ...
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

// TraverseInorder ...
func (n *Node) TraverseInorder() {
	if n.Left != nil {
		n.Left.TraverseInorder()
	}
	fmt.Println(*n.Data)
	if n.Right != nil {
		n.Right.TraverseInorder()
	}
}

// TraversePreorder ...
func (n *Node) TraversePreorder() {
	fmt.Println(*n.Data)
	if n.Left != nil {
		n.Left.TraversePreorder()
	}
	if n.Right != nil {
		n.Right.TraversePreorder()
	}
}

// TraversePostorder ...
func (n *Node) TraversePostorder() {
	if n.Left != nil {
		n.Left.TraversePostorder()
	}
	if n.Right != nil {
		n.Right.TraversePostorder()
	}
	fmt.Println(*n.Data)
}

// Lowest ...
func (n *Node) Lowest() int {
	if n.Left != nil {
		return n.Left.Lowest()
	}
	return *n.Data
}

// Highest ...
func (n *Node) Highest() int {
	if n.Right != nil {
		return n.Right.Highest()
	}
	return *n.Data
}

// Exist ...
func (n *Node) Exist(value int) bool {
	if *n.Data == value {
		return true
	}
	if *n.Data >= value {
		if n.Left != nil {
			fmt.Println("<-")
			return n.Left.Exist(value)
		}
		return false
	}
	if n.Right != nil {
		fmt.Println("->")
		return n.Right.Exist(value)
	}
	return false
}

// New ...
func New(elements []int) Node {
	tree := Node{}
	for _, element := range elements {
		tree.Insert(element)
	}
	return tree
}
