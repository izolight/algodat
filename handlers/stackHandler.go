package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/izolight/algodat/datastructures"
)

var stackTemplates = template.Must(template.ParseFiles("templates/header.html", "templates/footer.html", "templates/stack.html"))

type StackModel struct {
	Messages []string
	Errors   []string
	Stack    datastructures.Stack
	Peek     string
}

// View displays all values on the stack
func (s *Stack) View(w http.ResponseWriter, r *http.Request) {
	data := stackData{s.elements, ""}
	peek, err := s.peek()
	if err != nil {
		s.errors = append(s.errors, fmt.Sprintf("%v", err))
	} else {
		data.Peek = fmt.Sprintf("%d at position %d", peek, s.size-1)
	}
	view := viewData{"Stack", data, s.messages, s.errors}
	err = stackTemplates.ExecuteTemplate(w, "stack", view)
	if err != nil {
		fmt.Fprintf(w, "Could not render template: %v", err)
	}
	s.messages = s.messages[:0]
	s.errors = s.errors[:0]
}

// Push takes a value from a form and pushes it on the stack
func (s *Stack) Push(w http.ResponseWriter, r *http.Request) {
	new, err := strconv.Atoi(r.FormValue("new"))
	if err != nil {
		s.errors = append(s.errors, fmt.Sprintf("Could not parse form value: %v", err))
	} else {
		err = s.push(new)
		if err != nil {
			s.errors = append(s.errors, fmt.Sprintf("%v", err))
		} else {
			s.messages = append(s.messages, fmt.Sprintf("Pushed %d to stack", new))
		}
	}
	http.Redirect(w, r, "/stack", http.StatusSeeOther)
}

// Pop pops the top value from the stack
func (s *Stack) Pop(w http.ResponseWriter, r *http.Request) {
	popped, err := s.pop()
	if err != nil {
		s.errors = append(s.errors, fmt.Sprintf("%v", err))
	} else {
		s.messages = append(s.messages, fmt.Sprintf("Popped %d from stack", popped))
	}
	http.Redirect(w, r, "/stack", http.StatusSeeOther)
}
