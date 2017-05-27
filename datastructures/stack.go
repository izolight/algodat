package datastructures

import (
	"fmt"
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseFiles("templates/header.html", "templates/footer.html", "templates/viewstack.html"))

// Stack has the a pointer to the top element in the Stack, the current size and the maximum allowed size
type Stack struct {
	top     *element
	Size    int
	maxSize int
}

// element consists of the value and a pointer to the next element in the stack
type element struct {
	Value int
	next  *element
}

type viewData struct {
	Title string
	Data  interface{}
}

type stackData struct {
	Position, Value int
}

// NewStack creates a new stack with the specified maxSize
func NewStack(maxSize int) *Stack {
	return &Stack{nil, 0, maxSize}
}

func (s *Stack) Push(val int) error {
	if s.isFull() {
		return fmt.Errorf("Can't push to full stack. Size: %d", s.Size)
	}
	top := element{val, s.top}
	s.top = &top
	s.Size++
	return nil
}

func (s *Stack) Pop() (int, error) {
	if !s.isEmpty() {
		val := s.top.Value
		s.top = s.top.next
		s.Size--
		return val, nil
	}
	return 0, fmt.Errorf("Can't pop from empty stack")
}

// Checks whether the stack is empty
func (s *Stack) isEmpty() bool {
	if s.Size > 0 {
		return false
	}
	return true
}

func (s *Stack) isFull() bool {
	if s.Size == s.maxSize {
		return true
	}
	return false
}

// Returns the top element if any
func (s *Stack) getTop() (*element, error) {
	if s.isEmpty() {
		return nil, fmt.Errorf("Can't return top of empty stack")
	}
	return s.top, nil
}

// Returns the next element if any
func (e *element) Next() (*element, error) {
	if e.next != nil {
		return e.next, nil
	}
	return nil, fmt.Errorf("No more nodes")
}

func (s *Stack) GetAll(w http.ResponseWriter, r *http.Request) {
	e, err := s.getTop()
	if err != nil {
		fmt.Fprintf(w, "Could not get top of stack: %v", err)
	}
	var d []stackData
	pos := s.Size
	for {
		d = append(d, stackData{pos, e.Value})
		e, err = e.Next()
		if err != nil {
			break
		}
		pos--
	}
	data := viewData{"Stack", d}
	err = templates.ExecuteTemplate(w, "viewstack", data)
	if err != nil {
		fmt.Fprintf(w, "Could not render template: %v", err)
	}
}
