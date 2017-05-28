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

type stackData struct {
	Position, Value int
}

// Stack has the a pointer to the top element in the Stack, the current size and the maximum allowed size
type Stack struct {
	top     *stackElement
	size    int
	maxSize int
}

// stackElement consists of the value and a pointer to the next stackElement in the stack
type stackElement struct {
	Value int
	next  *stackElement
}

// NewStack creates a new stack with the specified maxSize
func NewStack(maxSize int) *Stack {
	return &Stack{nil, 0, maxSize}
}

// Pushes element on the stack
func (s *Stack) push(val int) error {
	if s.isFull() {
		return fmt.Errorf("Can't push to full stack. Size: %d", s.size)
	}
	top := stackElement{val, s.top}
	s.top = &top
	s.size++
	return nil
}

// Pops element from the stack and returns it
func (s *Stack) pop() (int, error) {
	if !s.isEmpty() {
		val := s.top.Value
		s.top = s.top.next
		s.size--
		return val, nil
	}
	return 0, fmt.Errorf("Can't pop from empty stack")
}

// Checks whether the stack is empty
func (s *Stack) isEmpty() bool {
	if s.size > 0 {
		return false
	}
	return true
}

func (s *Stack) isFull() bool {
	if s.size == s.maxSize {
		return true
	}
	return false
}

// Returns the top element if any
func (s *Stack) peek() (*stackElement, error) {
	if s.isEmpty() {
		return nil, fmt.Errorf("Can't return top of empty stack")
	}
	return s.top, nil
}

// Returns the next element if any
func (e *stackElement) Next() (*stackElement, error) {
	if e.next != nil {
		return e.next, nil
	}
	return nil, fmt.Errorf("No more nodes")
}

// View displays all values on the stack
func (s *Stack) View(w http.ResponseWriter, r *http.Request) {
	var d []stackData
	e, err := s.peek()
	if err != nil {
		errors = append(errors, fmt.Sprintf("Could not get top of stack: %v", err))
	} else {
		pos := s.size
		for {
			d = append(d, stackData{pos, e.Value})
			e, err = e.Next()
			if err != nil {
				break
			}
			pos--
		}
	}
	data := viewData{"Stack", d, errors}
	err = stackTemplates.ExecuteTemplate(w, "viewstack", data)
	if err != nil {
		fmt.Fprintf(w, "Could not render template: %v", err)
	}
	errors = errors[:0]
}

// Add takes a value from a form and pushes it on the stack
func (s *Stack) Add(w http.ResponseWriter, r *http.Request) {
	new, err := strconv.Atoi(r.FormValue("new"))
	if err != nil {
		errors = append(errors, fmt.Sprintf("Could not parse form value: %v", err))
	} else {
		err = s.push(new)
		if err != nil {
			errors = append(errors, fmt.Sprintf("Could not push to stack: %v", err))
		}
	}
	http.Redirect(w, r, "/stack", http.StatusSeeOther)
}

// Remove pops the top value from the stack
func (s *Stack) Remove(w http.ResponseWriter, r *http.Request) {
	_, err := s.pop()
	if err != nil {
		errors = append(errors, fmt.Sprintf("Could not pop from stack: %v", err))
	}
	http.Redirect(w, r, "/stack", http.StatusSeeOther)
}
