package tree

import (
	"fmt"
)

// Heap is a tree which the root / parent always bigger or smaller than the childern
type Heap struct {
	Type  string
	Items []int
}

func (h *Heap) length() int {
	return len(h.Items)
}

func (h *Heap) lastIndex() int {
	return h.length() - 1
}

func (h *Heap) getParentIndex(index int) int {
	return (index - 1) / 2
}

func (h *Heap) getLeftIndex(index int) int {
	return (2 * index) + 1
}

func (h *Heap) getRightIndex(index int) int {
	return (2 * index) + 2
}

func (h *Heap) getParentValue(index int) int {
	return h.Items[h.getParentIndex(index)]
}

func (h *Heap) getLeftValue(index int) int {
	return h.Items[h.getLeftIndex(index)]
}

func (h *Heap) getRightValue(index int) int {
	return h.Items[h.getRightIndex(index)]
}

// swap tree position
func (h *Heap) swap(indexA, indexB int) {
	temp := h.Items[indexA]
	h.Items[indexA] = h.Items[indexB]
	h.Items[indexB] = temp
}

func (h *Heap) hasLeft(index int) bool {
	return h.getLeftIndex(index) <= h.lastIndex()
}

func (h *Heap) hasRight(index int) bool {
	return h.getRightIndex(index) <= h.lastIndex()
}

func (h *Heap) hasParent(index int) bool {
	return index > 0
}

// Insert a new element into the heap and readjust the position accordingly
func (h *Heap) Insert(value int) {
	// add new value to tail
	h.Items = append(h.Items, value)

	// reorder node bottom top
	h.heapifyUp()
}

func (h *Heap) compareParent(index int) bool {
	if h.Type == "min" {
		return h.getParentValue(index) > h.Items[index]
	}
	return h.getParentValue(index) < h.Items[index]
}

// reorder the last element position
func (h *Heap) heapifyUp() {
	index := h.lastIndex()
	for h.hasParent(index) && h.compareParent(index) {
		parentIndex := h.getParentIndex(index)
		h.swap(parentIndex, index)
		index = parentIndex
	}
}

func (h *Heap) compareSibling(index int) bool {
	if h.Type == "min" {
		return h.getRightValue(index) < h.getLeftValue(index)
	}

	return h.getRightValue(index) > h.getLeftValue(index)
}

// reorder root to correct position
func (h *Heap) heapifyDown() {
	index := 0

	// check if current has child, no need to check both
	for h.hasLeft(index) {
		childIndex := h.getLeftIndex(index)

		// swap smaller index if right is smaller than left child
		if h.hasRight(index) && h.compareSibling(index) {
			childIndex = h.getRightIndex(index)
		}

		// stop iteration if current is smaller than child value
		if h.Items[index] < h.Items[childIndex] {
			break
		}

		// swap element position and continue iteration
		h.swap(index, childIndex)
		index = childIndex
	}
}

// DeleteRoot removes top element and readjust new inherit
func (h *Heap) DeleteRoot() {
	if h.lastIndex() == 0 {
		return
	}

	// replace root with furthest tail then remove tail
	h.Items[0] = h.Items[h.lastIndex()]
	h.Items = h.Items[:h.lastIndex()]

	// reorder heap top bottom
	h.heapifyDown()
}

// Print every element in heap
func (h *Heap) Print() {
	fmt.Println(h.Items)
}

// NewHeap returns new heap instance (min, max)
func NewHeap(elements []int, kind string) Heap {
	heap := Heap{
		Items: []int{},
		Type:  kind,
	}

	for _, element := range elements {
		heap.Insert(element)
	}

	return heap
}
