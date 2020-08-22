package linkedlist

import (
	"fmt"
)

// Node represent an element in linkedlist
type Node struct {
	Next *Node
	Prev *Node
	Data int
}

// LinkedList represent the whole list
type LinkedList struct {
	Head *Node
}

// Append add element to the end of linkedlist
func (l *LinkedList) Append(val int) {
	// check if head is exist
	if l.Head == nil {
		l.Head = &Node{
			Data: val,
		}
		return
	}

	// loop until the end of linkedlist
	current := l.Head
	for current.Next != nil {
		current = current.Next
	}

	// connect tail to new node
	current.Next = &Node{
		Data: val,
	}
}

// Prepend add element to the begining of the linkedlist
func (l *LinkedList) Prepend(val int) {
	// check if head is exist
	if l.Head == nil {
		l.Head = &Node{
			Data: val,
		}
		return
	}

	// connect head to the new node
	currentHead := l.Head
	newHead := &Node{
		Data: val,
		Next: currentHead,
	}
	l.Head = newHead
}

// Delete a value within the linkedlist
func (l *LinkedList) Delete(val int) {
	// check if head is exist
	if l.Head == nil {
		return
	}

	// if element that will be deleted is head, delegate head to next element
	if l.Head.Data == val {
		l.Head = l.Head.Next
		return
	}

	// loop until element is found, then removes previos and next connection of deleted element
	current := l.Head
	for current.Next != nil {
		if current.Next.Data == val {
			current.Next = current.Next.Next
			return
		}
		current = current.Next
	}

}

// Print every element in linkedlist
func (l *LinkedList) Print() {
	current := l.Head
	for current != nil {
		fmt.Println(current.Data)
		current = current.Next
	}
}

// New returns new Linkedlist instance
func New(elements []int) LinkedList {
	ll := LinkedList{}

	for _, item := range elements {
		ll.Append(item)
	}

	return ll
}
