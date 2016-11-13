package queue

import (
	"fmt"
)

type Queue struct {
	front, rear *Node
	size        int
}

type Node struct {
	next *Node
	val  int
}

func Enqueue(q *Queue, val int) {
	rear := Node{nil, val}
	if q.size == 0 {
		q.front = &rear
	} else {
		q.rear.next = &rear
	}
	q.rear = &rear
	q.size++
}

func Dequeue(q *Queue) int {
	if q.size > 0 {
		front := q.front.val
		q.front = q.front.next
		q.size--
		return front
	}
	return 0
}

func Front(q *Queue) int {
	return q.front.val
}
func Rear(q *Queue) int {
	return q.rear.val
}

func Size(q *Queue) int {
	return q.size
}

func IsEmpty(q *Queue) bool {
	if q.size > 0 {
		return false
	}
	return true
}

func PrintQueue(q *Queue) {
	if !IsEmpty(q) {
		fmt.Printf("Front: %d\tRear: %d\tSize: %d\tEmpty: %t\n", Front(q), Rear(q), Size(q), IsEmpty(q))
	} else {
		fmt.Printf("Size: %d\tEmpty: %t\n", Size(q), IsEmpty(q))
	}
}