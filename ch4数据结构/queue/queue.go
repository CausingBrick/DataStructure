// package queue is the queue example made up by slice

package queue

// Queue is the common type for queue.
type Queue []interface{}

// Enqueue push element to queue
func (q *Queue) Enqueue(x ...interface{}) {
	for _, val := range x {
		*q = append(*q, val)
	}
}

// Dequeue pop element for queue
func (q *Queue) Dequeue() interface{} {
	queue := *q
	head := queue[0]
	*q = queue[1:]
	return head
}

// IsEmpty return if the queue is empty
func (q *Queue) IsEmpty() bool {
	queue := *q
	return len(queue) == 0
}
