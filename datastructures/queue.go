package datastructures

import "html/template"
import "fmt"

var queueTemplates = template.Must(template.ParseFiles("templates/header.html", "templates/footer.html", "templates/queue.html"))

type Queue struct {
	size, maxSize int
	elements      []int
}

func NewQueue(maxSize int) *Queue {
	return &Queue{0, maxSize, nil}
}

func (q *Queue) enqueue(val int) error {
	if q.isFull() {
		return fmt.Errorf("Can't enqueue to full queue")
	}
	q.elements = append([]int{val}, q.elements...)
	q.size++
	return nil
}

func (q *Queue) dequeue() error {
	if q.isEmpty() {
		return fmt.Errorf("Cant't dequeue from empty queue")
	}
	q.elements = q.elements[:len(q.elements)-1]
	q.size--
	return nil
}

func (q *Queue) peek() (int, error) {
	if q.isEmpty() {
		return 0, fmt.Errorf("Cant't peek into empty queue")
	}
	return q.elements[len(q.elements)-1], nil
}

func (q *Queue) isFull() bool {
	if q.size == q.maxSize {
		return true
	}
	return false
}

func (q *Queue) isEmpty() bool {
	if q.size == 0 {
		return true
	}
	return false
}
