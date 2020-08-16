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

func recursiveQuickSort(arr []int, left int, right int) {
	// bounce check
	if left >= right {
		return
	}

	pivot := arr[(left+right)/2]
	index := partition(arr, left, right, pivot)

	recursiveQuickSort(arr, left, index-1)
	recursiveQuickSort(arr, index+1, right)
}

func partition(arr []int, left int, right int, pivot int) int {
	for left <= right {
		for arr[left] < pivot {
			left++
		}

		for arr[right] > pivot {
			right--
		}

		if left <= right {
			swap(&arr, left, right)
			left++
			right--
		}
	}
	return left
}

func swap(l *[]int, i, j int) {
	temp := (*l)[i]
	(*l)[i] = (*l)[i+1]
	(*l)[i+1] = temp
}
