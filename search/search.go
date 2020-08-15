package search

import (
	"errors"
	"fmt"
)

// List ...
type List []int

// LinearSearch O(n)
func (l *List) LinearSearch(value int) (int, error) {
	for index, val := range *l {
		if value == val {
			fmt.Println(index)
			return index, nil
		}
	}
	return 0, errors.New("value not found")
}

// BinarySearch O(Log(n))
func (l *List) BinarySearch(value int) {
	hasValue := recursiveBinarySearch(*l, value, 0, len(*l)-1)
	fmt.Println(hasValue)
}

func recursiveBinarySearch(arr []int, value int, leftBound int, rightBound int) bool {
	// if bound overlapping, data not found
	if leftBound > rightBound {
		return false
	}

	midBound := (leftBound + rightBound) / 2
	midValue := arr[midBound]

	if value == midValue {
		return true
	} else if value < midValue {
		return recursiveBinarySearch(arr, value, leftBound, midBound-1)
	}

	return recursiveBinarySearch(arr, value, midBound+1, rightBound)
}
