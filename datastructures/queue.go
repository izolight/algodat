package datastructures

import (
	"fmt"
)

// Queue represents the queue datastructure
type Queue struct {
	maxSize  int
	elements []fmt.Stringer
}

// Empty checks whether the queue is empty
func (q *Queue) Empty() bool {
	if len(q.elements) == 0 {
		return true
	}
	return false
}

// Size returns the number of elements in the queue
func (q *Queue) Size() int {
	return len(q.elements)
}

// Elements returns all elements in the queue
func (q *Queue) Elements() []fmt.Stringer {
	return q.elements
}

// Full checks whether the queue is full
func (q *Queue) Full() bool {
	if q.Size() == q.maxSize {
		return true
	}
	return false
}

// NewQueue returns a queue with the specified maxSize
func NewQueue(maxSize int) *Queue {
	return &Queue{maxSize, nil}
}

// Enqueue adds an element to the end of the queue
func (q *Queue) Enqueue(val fmt.Stringer) error {
	if q.Full() {
		return fmt.Errorf("Can't enqueue to full queue. Size: %d", q.Size())
	}
	q.elements = append(q.elements, val)
	return nil
}

// Dequeue removes an element from the front of the Queue
func (q *Queue) Dequeue() (fmt.Stringer, error) {
	if q.Empty() {
		return nil, fmt.Errorf("Cant't dequeue from empty queue")
	}
	dequeued := q.elements[0]
	q.elements = q.elements[1:]
	return dequeued, nil
}

// Front returns the first element in the queue without removing it
func (q *Queue) Front() (fmt.Stringer, error) {
	if q.Empty() {
		return nil, fmt.Errorf("Cant't peek into empty queue")
	}
	return q.elements[0], nil
}

// Back returns the last element in the queue without removing it
func (q *Queue) Back() (fmt.Stringer, error) {
	if q.Empty() {
		return nil, fmt.Errorf("Cant't peek into empty queue")
	}
	return q.elements[q.Size()-1], nil
}
