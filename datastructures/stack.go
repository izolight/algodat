package datastructures

import (
	"fmt"
)

// Stack has the a pointer to the top element in the Stack, the current size and the maximum allowed size
type Stack struct {
	maxSize  int
	elements []fmt.Stringer
}

// Empty checks whether the stack is empty
func (s *Stack) Empty() bool {
	if len(s.elements) == 0 {
		return true
	}
	return false
}

// Size returns the number of elements in the stack
func (s *Stack) Size() int {
	return len(s.elements)
}

// Elements returns all elements in the stack
func (s *Stack) Elements() []fmt.Stringer {
	return s.elements
}

// Full checks whether the stack is full
func (s *Stack) Full() bool {
	if s.Size() == s.maxSize {
		return true
	}
	return false
}

// NewStack creates a new stack with the specified maxSize
func NewStack(maxSize int) *Stack {
	return &Stack{maxSize, nil}
}

// Top returns the top element if any
func (s *Stack) Top() (fmt.Stringer, error) {
	if s.Empty() {
		return nil, fmt.Errorf("Can't get top of empty stack")
	}
	return s.elements[len(s.elements)-1], nil
}

// Push adds an element on top of the stack
func (s *Stack) Push(val fmt.Stringer) error {
	if s.Full() {
		return fmt.Errorf("Can't push to full stack. Size: %d", len(s.elements))
	}
	s.elements = append(s.elements, val)
	return nil
}

// Pop removes and returns the top element of the stack
func (s *Stack) Pop() (fmt.Stringer, error) {
	if s.Empty() {
		return nil, fmt.Errorf("Can't pop from empty stack")
	}
	popped := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return popped, nil
}
