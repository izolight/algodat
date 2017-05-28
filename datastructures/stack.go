package datastructures

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

var stackTemplates = template.Must(template.ParseFiles("templates/header.html", "templates/footer.html", "templates/stack.html"))
var errors []string

type viewData struct {
	Title  string
	Data   interface{}
	Errors []string
}

// Stack has the a pointer to the top element in the Stack, the current size and the maximum allowed size
type Stack struct {
	size, maxSize int
	elements      []int
}

// NewStack creates a new stack with the specified maxSize
func NewStack(maxSize int) *Stack {
	return &Stack{0, maxSize, nil}
}

// Pushes element on the stack
func (s *Stack) push(val int) error {
	if s.isFull() {
		return fmt.Errorf("Can't push to full stack. Size: %d", s.size)
	}
	s.elements = append(s.elements, val)
	s.size++
	return nil
}

// Pops element from the stack and returns it
func (s *Stack) pop() error {
	if s.isEmpty() {
		return fmt.Errorf("Can't pop from empty stack")
	}
	s.elements = s.elements[:len(s.elements)-1]
	s.size--
	return nil
}

// Checks whether the stack is empty
func (s *Stack) isEmpty() bool {
	if s.size == 0 {
		return true
	}
	return false
}

func (s *Stack) isFull() bool {
	if s.size == s.maxSize {
		return true
	}
	return false
}

// Returns the top element if any
func (s *Stack) peek() (int, error) {
	if s.isEmpty() {
		return 0, fmt.Errorf("Can't return top of empty stack")
	}
	return s.elements[len(s.elements)-1], nil
}

// View displays all values on the stack
func (s *Stack) View(w http.ResponseWriter, r *http.Request) {
	data := viewData{"Stack", s.elements, errors}
	err := stackTemplates.ExecuteTemplate(w, "stack", data)
	if err != nil {
		fmt.Fprintf(w, "Could not render template: %v", err)
	}
	errors = errors[:0]
}

// Push takes a value from a form and pushes it on the stack
func (s *Stack) Push(w http.ResponseWriter, r *http.Request) {
	new, err := strconv.Atoi(r.FormValue("new"))
	if err != nil {
		errors = append(errors, fmt.Sprintf("Could not parse form value: %v", err))
	} else {
		err = s.push(new)
		if err != nil {
			errors = append(errors, fmt.Sprintf("%v", err))
		}
	}
	http.Redirect(w, r, "/stack", http.StatusSeeOther)
}

// Pop pops the top value from the stack
func (s *Stack) Pop(w http.ResponseWriter, r *http.Request) {
	err := s.pop()
	if err != nil {
		errors = append(errors, fmt.Sprintf("%v", err))
	}
	http.Redirect(w, r, "/stack", http.StatusSeeOther)
}
