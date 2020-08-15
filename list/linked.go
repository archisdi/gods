package list

import (
	"fmt"
)

// Node ...
type Node struct {
	Data int
	Next *Node
}

// LinkedQueue ...
type LinkedQueue struct {
	Head *Node
	Tail *Node
}

// IsEmpty ...
func (lq *LinkedQueue) IsEmpty() bool {
	return lq.Head == nil
}

// Peek ...
func (lq *LinkedQueue) Peek() int {
	return lq.Head.Data
}

// Enqueue ...
func (lq *LinkedQueue) Enqueue(val int) {
	node := &Node{
		Data: val,
	}

	if lq.Head == nil {
		lq.Head = node
		lq.Tail = node
		return
	}

	currentTail := lq.Tail
	currentTail.Next = node
	lq.Tail = node
}

// Dequeue ...
func (lq *LinkedQueue) Dequeue() {
	if lq.Head == nil {
		return
	}

	if lq.Head == lq.Tail {
		lq.Head = nil
		lq.Tail = nil
		return
	}

	lq.Head = lq.Head.Next
}

// Print ...
func (lq *LinkedQueue) Print() {
	if lq.Head == nil {
		return
	}

	current := lq.Head
	for current != nil {
		fmt.Println(current.Data)
		current = current.Next
	}
}

// LinkedStack ...
type LinkedStack struct {
	Top *Node
}

// IsEmpty ...
func (ls *LinkedStack) IsEmpty() bool {
	return ls.Top == nil
}

// Peek ...
func (ls *LinkedStack) Peek() int {
	return ls.Top.Data
}

// Add ...
func (ls *LinkedStack) Add(val int) {
	node := &Node{
		Data: val,
	}

	if ls.Top == nil {
		ls.Top = node
		return
	}

	node.Next = ls.Top
	ls.Top = node
}

// Pop ...
func (ls *LinkedStack) Pop() {
	if ls.Top == nil {
		return
	}
	ls.Top = ls.Top.Next
}

// Print ...
func (ls *LinkedStack) Print() {
	if ls.Top == nil {
		return
	}

	current := ls.Top
	for current != nil {
		fmt.Println(current.Data)
		current = current.Next
	}
}

// NewLinkedStack ...
func NewLinkedStack(elements []int) LinkedStack {
	ls := LinkedStack{}

	for _, element := range elements {
		ls.Add(element)
	}

	return ls
}

// NewLinkedQueue ...
func NewLinkedQueue(elements []int) LinkedQueue {
	lq := LinkedQueue{}

	for _, element := range elements {
		lq.Enqueue(element)
	}

	return lq
}
