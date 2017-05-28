package datastructures

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

var queueTemplates = template.Must(template.ParseFiles("templates/header.html", "templates/footer.html", "templates/queue.html"))

// Queue represents the queue datastructure
type Queue struct {
	size, maxSize int
	elements      []int
	messages      []string
	errors        []string
}

type queueData struct {
	Elements []int
	Peek     string
}

// NewQueue returns a queue with the specified maxSize
func NewQueue(maxSize int) *Queue {
	return &Queue{0, maxSize, nil, nil, nil}
}

func (q *Queue) enqueue(val int) error {
	if q.isFull() {
		return fmt.Errorf("Can't enqueue to full queue. Size: %d", q.size)
	}
	q.elements = append(q.elements, val)
	q.size++
	return nil
}

func (q *Queue) dequeue() (int, error) {
	if q.isEmpty() {
		return 0, fmt.Errorf("Cant't dequeue from empty queue")
	}
	dequeued := q.elements[0]
	q.elements = q.elements[1:]
	q.size--
	return dequeued, nil
}

func (q *Queue) peek() (int, error) {
	if q.isEmpty() {
		return 0, fmt.Errorf("Cant't peek into empty queue")
	}
	return q.elements[0], nil
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
	data := queueData{q.elements, ""}
	peek, err := q.peek()
	if err != nil {
		q.errors = append(q.errors, fmt.Sprintf("%v", err))
	} else {
		data.Peek = fmt.Sprintf("%d at position %d", peek, 0)
	}
	view := viewData{"Queue", data, q.messages, q.errors}
	err = queueTemplates.ExecuteTemplate(w, "queue", view)
	if err != nil {
		fmt.Fprintf(w, "Could not render template: %v", err)
	}
	q.messages = q.messages[:0]
	q.errors = q.errors[:0]
}

// Enqueue takes the form value and enqueues it
func (q *Queue) Enqueue(w http.ResponseWriter, r *http.Request) {
	new, err := strconv.Atoi(r.FormValue("new"))
	if err != nil {
		q.errors = append(q.errors, fmt.Sprintf("Could not parse form value: %v", err))
	} else {
		err = q.enqueue(new)
		if err != nil {
			q.errors = append(q.errors, fmt.Sprintf("%v", err))
		} else {
			q.messages = append(q.messages, fmt.Sprintf("Enqueued %d to Queue", new))
		}
	}
	http.Redirect(w, r, "/queue", http.StatusSeeOther)
}

// Dequeue dequeues from the queue
func (q *Queue) Dequeue(w http.ResponseWriter, r *http.Request) {
	dequeued, err := q.dequeue()
	if err != nil {
		q.errors = append(q.errors, fmt.Sprintf("%v", err))
	} else {
		q.messages = append(q.messages, fmt.Sprintf("Dequeued %d from Queue", dequeued))
	}
	http.Redirect(w, r, "/queue", http.StatusSeeOther)
}
