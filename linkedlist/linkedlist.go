package linkedlist

import (
	"fmt"
)

// Node ...
type Node struct {
	Next *Node
	Prev *Node
	Data int
}

// LinkedList ...
type LinkedList struct {
	Head *Node
}

// Append ..
func (l *LinkedList) Append(val int) {
	if l.Head == nil {
		l.Head = &Node{
			Data: val,
		}
		return
	}

	current := l.Head
	for current.Next != nil {
		current = current.Next
	}

	current.Next = &Node{
		Data: val,
	}
}

// Prepend ...
func (l *LinkedList) Prepend(val int) {
	if l.Head == nil {
		l.Head = &Node{
			Data: val,
		}
		return
	}

	currentHead := l.Head
	newHead := &Node{
		Data: val,
		Next: currentHead,
	}
	l.Head = newHead
}

// Delete ...
func (l *LinkedList) Delete(val int) {
	if l.Head == nil {
		return
	}

	if l.Head.Data == val {
		l.Head = l.Head.Next
		return
	}

	current := l.Head
	for current.Next != nil {
		if current.Next.Data == val {
			current.Next = current.Next.Next
			return
		}
		current = current.Next
	}

}

// Print ...
func (l *LinkedList) Print() {
	current := l.Head
	for current != nil {
		fmt.Println(current.Data)
		current = current.Next
	}
}

// New ...
func New(elements []int) LinkedList {
	ll := LinkedList{}

	for _, item := range elements {
		ll.Append(item)
	}

	return ll
}
