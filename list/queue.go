package list

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
