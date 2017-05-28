package datastructures

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

var stackTemplates = template.Must(template.ParseFiles("templates/header.html", "templates/footer.html", "templates/stack.html"))

// Stack has the a pointer to the top element in the Stack, the current size and the maximum allowed size
type Stack struct {
	size, maxSize int
	elements      []int
	messages      []string
	errors        []string
}

type stackData struct {
	Elements []int
	Peek     string
}

// NewStack creates a new stack with the specified maxSize
func NewStack(maxSize int) *Stack {
	return &Stack{0, maxSize, nil, nil, nil}
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
func (s *Stack) pop() (int, error) {
	if s.isEmpty() {
		return 0, fmt.Errorf("Can't pop from empty stack")
	}
	popped := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	s.size--
	return popped, nil
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
