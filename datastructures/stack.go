package datastructures

import (
	"fmt"
)

// Stack has the a pointer to the top element in the Stack, the current size and the maximum allowed size
type Stack struct {
	top     *Element
	Size    int
	maxSize int
}

// Element consists of the value and a pointer to the next Element in the stack
type Element struct {
	Value int
	next  *Element
}

// Creates a new stack with the specified maxSize
func NewStack(maxSize int) *Stack {
	return &Stack{nil, 0, maxSize}
}

func (s *Stack) Push(val int) error {
	if s.isFull() {
		return fmt.Errorf("Can't push to full stack. Size: %d", s.Size)
	}
	top := Element{val, s.top}
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
func (s *Stack) Top() (*Element, error) {
	if s.isEmpty() {
		return nil, fmt.Errorf("Can't return top of empty stack")
	}
	return s.top, nil
}

// Returns the next element if any
func (e *Element) Next() (*Element, error) {
	if e.next != nil {
		return e.next, nil
	}
	return nil, fmt.Errorf("No more nodes")
}