package datastructures

import (
	"fmt"
)

// Stack has the a pointer to the top element in the Stack, the current size and the maximum allowed size
type Stack struct {
	maxSize  int
	elements []fmt.Stringer
}

// NewStack creates a new stack with the specified maxSize
func NewStack(maxSize int) *Stack {
	return &Stack{maxSize, nil}
}

// Pushes element on the stack
func (s *Stack) push(val fmt.Stringer) error {
	if s.isFull() {
		return fmt.Errorf("Can't push to full stack. Size: %d", len(s.elements))
	}
	s.elements = append(s.elements, val)
	return nil
}

// Pops element from the stack and returns it
func (s *Stack) pop() (string, error) {
	if s.isEmpty() {
		return "", fmt.Errorf("Can't pop from empty stack")
	}
	popped := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return popped.String(), nil
}

// Checks whether the stack is empty
func (s *Stack) isEmpty() bool {
	if len(s.elements) == 0 {
		return true
	}
	return false
}

func (s *Stack) isFull() bool {
	if len(s.elements) == s.maxSize {
		return true
	}
	return false
}

// Returns the top element if any
func (s *Stack) peek() (string, error) {
	if s.isEmpty() {
		return "", fmt.Errorf("Can't peek into empty stack")
	}
	return s.elements[len(s.elements)-1].String(), nil
}
