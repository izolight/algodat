package main

import (
	"time"
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

func enqueue(q *Queue, val int) {
	rear := Node{nil, val}
	if q.size == 0 {
		q.front = &rear
	} else {
		q.rear.next = &rear
	}
	q.rear = &rear
	q.size++
}

func dequeue(q *Queue) int {
	if q.size > 0 {
		front := q.front.val
		q.front = q.front.next
		q.size--
		return front
	}
	return 0
}

func front(q *Queue) int {
	return q.front.val
}
func rear(q *Queue) int {
	return q.rear.val
}

func queueSize(q *Queue) int {
	return q.size
}

func queueIsEmpty(q *Queue) bool {
	if q.size > 0 {
		return false
	}
	return true
}

func printQueue(q *Queue) {
	if !queueIsEmpty(q) {
		fmt.Printf("Front: %d\tRear: %d\tSize: %d\tEmpty: %t\n", front(q), rear(q), queueSize(q), queueIsEmpty(q))
	} else {
		fmt.Printf("Size: %d\tEmpty: %t\n", queueSize(q), queueIsEmpty(q))
	}
}

func main() {
	myQueue := new(Queue)
	const max int = 16 * 1024 * 1024 + 2
	start := time.Now()
	for i := 0; i < max; i++ {
		enqueue(myQueue, i)
	}
	end := time.Now()
	fmt.Printf("Time taken to store %d elements: %v\n", max, end.Sub(start))
	printQueue(myQueue)
	start = time.Now()
	for i := 0; i < max / 2; i++ {
		dequeue(myQueue)
	}
	end = time.Now()
	fmt.Printf("Time taken to remove %d elements: %v\n", max / 2, end.Sub(start))
	printQueue(myQueue)
}
