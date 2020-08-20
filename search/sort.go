package search

import (
	"fmt"
)

func (l *List) swap(i, j int) {
	temp := (*l)[i]
	(*l)[i] = (*l)[j]
	(*l)[j] = temp
}

func (l *List) length() int {
	return len(*l)
}

// Print ...
func (l *List) Print() {
	fmt.Println(*l)
}

// BubbleSort ...
func (l *List) BubbleSort() {
	arrLength := l.length() - 1
	isSorted := false

	for !isSorted {
		isSorted = true
		for i := 0; i < arrLength; i++ {
			if (*l)[i] > (*l)[i+1] {
				l.swap(i, i+1)
				isSorted = false
			}
		}
		arrLength--
	}
}

// QuickSort ...
func (l *List) QuickSort() {
	recursiveQuickSort(*l, 0, l.length()-1)
}

func recursiveQuickSort(arr List, start int, end int) {
	// bounce check
	if start >= end {
		return
	}
	index := partition(arr, start, end)
	recursiveQuickSort(arr, start, index-1)
	recursiveQuickSort(arr, index+1, end)
}

func partition(arr List, start int, end int) int {
	pivot := arr[end]
	pIndex := start
	for i := start; i < end; i++ {
		if arr[i] <= pivot {
			arr.swap(i, pIndex)
			pIndex++
		}
	}
	arr.swap(pIndex, end)
	return pIndex
}

// MergeSort ..
func (l *List) MergeSort() {
	recursiveMergeSort(*l, 0, l.length()-1)
}

func recursiveMergeSort(arr List, start int, end int) {
	if start >= end {
		return
	}

	mid := (start + end) / 2

	recursiveMergeSort(arr, start, mid-1)
	recursiveMergeSort(arr, mid+1, end)
}

func mergeHalves(arr List, start int, end int) {

}
