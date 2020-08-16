package search

func (l *List) swap(i, j int) {
	temp := (*l)[i]
	(*l)[i] = (*l)[i+1]
	(*l)[i+1] = temp
}

func (l *List) length() int {
	return len(*l)
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

func recursiveQuickSort(arr []int, start int, end int) {
	// bounce check
	if start >= end {
		return
	}
	index := partition(arr, start, end)
	recursiveQuickSort(arr, start, index-1)
	recursiveQuickSort(arr, index+1, end)
}

func partition(arr []int, start int, end int) int {
	pivot := arr[end]
	pIndex := start
	for i := start; i < end; i++ {
		if arr[i] <= pivot {
			swap(&arr, i, pIndex)
			pIndex++
		}
	}
	swap(&arr, pIndex, end)
	return pIndex
}

func swap(l *[]int, i, j int) {
	temp := (*l)[i]
	(*l)[i] = (*l)[j]
	(*l)[j] = temp
}
