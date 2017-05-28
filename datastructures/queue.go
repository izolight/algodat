package datastructures

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

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
		return fmt.Errorf("Can't enqueue to full queue. Size: %d", q.size)
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

// View displays all values in the queue
func (q *Queue) View(w http.ResponseWriter, r *http.Request) {
	data := viewData{"Queue", q.elements, errors}
	err := queueTemplates.ExecuteTemplate(w, "queue", data)
	if err != nil {
		fmt.Fprintf(w, "Could not render template: %v", err)
	}
	errors = errors[:0]
}

// Enqueue takes the form value and enqueues it
func (q *Queue) Enqueue(w http.ResponseWriter, r *http.Request) {
	new, err := strconv.Atoi(r.FormValue("new"))
	if err != nil {
		errors = append(errors, fmt.Sprintf("Could not parse form value: %v", err))
	} else {
		err = q.enqueue(new)
		if err != nil {
			errors = append(errors, fmt.Sprintf("%v", err))
		}
	}
	http.Redirect(w, r, "/queue", http.StatusSeeOther)
}

// Dequeue dequeues from the queue
func (q *Queue) Dequeue(w http.ResponseWriter, r *http.Request) {
	err := q.dequeue()
	if err != nil {
		errors = append(errors, fmt.Sprintf("%v", err))
	}
	http.Redirect(w, r, "/queue", http.StatusSeeOther)
}
