package tree

import (
	"fmt"
)

// MinHeap ...
type MinHeap struct {
	Items []int
}

func (mh *MinHeap) length() int {
	return len(mh.Items)
}

func (mh *MinHeap) lastIndex() int {
	return mh.length() - 1
}

func (mh *MinHeap) getParentIndex(index int) int {
	return (index - 1) / 2
}

func (mh *MinHeap) getLeftIndex(index int) int {
	return (2 * index) + 1
}

func (mh *MinHeap) getRightIndex(index int) int {
	return (2 * index) + 2
}

func (mh *MinHeap) getParentValue(index int) int {
	return mh.Items[mh.getParentIndex(index)]
}

func (mh *MinHeap) getLeftValue(index int) int {
	return mh.Items[mh.getLeftIndex(index)]
}

func (mh *MinHeap) getRightValue(index int) int {
	return mh.Items[mh.getRightIndex(index)]
}

func (mh *MinHeap) swap(indexA, indexB int) {
	temp := mh.Items[indexA]
	mh.Items[indexA] = mh.Items[indexB]
	mh.Items[indexB] = temp
}

func (mh *MinHeap) hasLeft(index int) bool {
	return mh.getLeftIndex(index) <= mh.lastIndex()
}

func (mh *MinHeap) hasRight(index int) bool {
	return mh.getRightIndex(index) <= mh.lastIndex()
}

func (mh *MinHeap) hasParent(index int) bool {
	return index > 0
}

// Insert ...
func (mh *MinHeap) Insert(value int) {
	// add new value to tail
	mh.Items = append(mh.Items, value)

	// reorder node bottom top
	mh.heapifyUp()
}

func (mh *MinHeap) heapifyUp() {
	index := mh.lastIndex()
	for mh.hasParent(index) && mh.getParentValue(index) > mh.Items[index] {
		parentIndex := mh.getParentIndex(index)
		mh.swap(parentIndex, index)
		index = parentIndex
	}
}

func (mh *MinHeap) heapifyDown() {
	index := 0

	// check if current has child, no need to check both
	for mh.hasLeft(index) {
		childIndex := mh.getLeftIndex(index)

		// swap smaller index if right is smaller than left child
		if mh.hasRight(index) && mh.getRightValue(index) < mh.getLeftValue(index) {
			childIndex = mh.getRightIndex(index)
		}

		// stop iteration if current is smaller than child value
		if mh.Items[index] < mh.Items[childIndex] {
			break
		}

		// swap element position and continue iteration
		mh.swap(index, childIndex)
		index = childIndex
	}
}

// DeleteRoot ...
func (mh *MinHeap) DeleteRoot() {
	if mh.lastIndex() == 0 {
		return
	}
	// replace root with furthest tail then remove tail
	mh.Items[0] = mh.Items[mh.lastIndex()]
	mh.Items = mh.Items[:mh.lastIndex()]

	// reorder heap top bottom
	mh.heapifyDown()
}

// Print ...
func (mh *MinHeap) Print() {
	fmt.Println(mh.Items)
}

// NewMinHeap ...
func NewMinHeap(elements []int) MinHeap {
	heap := MinHeap{elements}
	return heap
}
