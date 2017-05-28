package datastructures

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

var templates = template.Must(template.ParseFiles("templates/header.html", "templates/footer.html", "templates/viewstack.html"))
var errors []string

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
	Title  string
	Data   interface{}
	Errors []string
}

type stackData struct {
	Position, Value int
}

// NewStack creates a new stack with the specified maxSize
func NewStack(maxSize int) *Stack {
	return &Stack{nil, 0, maxSize}
}

// Pushes element on the stack
func (s *Stack) Push(val int) error {
	if s.isFull() {
		return fmt.Errorf("Can't push to full stack. Size: %d", s.Size)
	}
	top := element{val, s.top}
	s.top = &top
	s.Size++
	return nil
}

// Pops element from the stack and returns it
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

// GetAll displays all values on the stack
func (s *Stack) View(w http.ResponseWriter, r *http.Request) {
	var d []stackData
	e, err := s.getTop()
	if err != nil {
		errors = append(errors, fmt.Sprintf("Could not get top of stack: %v", err))
	} else {
		pos := s.Size
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
	err = templates.ExecuteTemplate(w, "viewstack", data)
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
		err = s.Push(new)
		if err != nil {
			errors = append(errors, fmt.Sprintf("Could not push to stack: %v", err))
		}
	}
	http.Redirect(w, r, "/stack", http.StatusSeeOther)
}

// Remove pops the top value from the stack
func (s *Stack) Remove(w http.ResponseWriter, r *http.Request) {
	_, err := s.Pop()
	if err != nil {
		errors = append(errors, fmt.Sprintf("Could not pop from stack: %v", err))
	}
	http.Redirect(w, r, "/stack", http.StatusSeeOther)
}
