package list

// Stack ...
type Stack []interface{}

// Push ...
func (s *Stack) Push(value interface{}) {
	*s = append(*s, value)
}

// Pop ...
func (s *Stack) Pop() interface{} {
	idx := len(*s) - 1
	value := (*s)[idx]
	*s = (*s)[:idx]
	return value
}

// IsEmpty ...
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Contains ...
func (s *Stack) Contains(value interface{}) (int, bool) {
	for index, item := range *s {
		if item == value {
			return index, true
		}
	}
	return 0, false
}

// Queue ...
type Queue []interface{}

// Enqueue ...
func (q *Queue) Enqueue(value interface{}) {
	*q = append([]interface{}{value}, *q...)
}

// Dequeue ...
func (q *Queue) Dequeue() interface{} {
	value := (*q)[0]
	*q = (*q)[1:]
	return value
}

// IsEmpty ...
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

// Contains ...
func (q *Queue) Contains(value interface{}) bool {
	for _, item := range *q {
		if item == value {
			return true
		}
	}
	return false
}
