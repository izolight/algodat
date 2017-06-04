package handlers

import (
	"fmt"
	"net/http"
	"strconv"
)

type queueData struct {
	Elements []int
	Peek     string
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
