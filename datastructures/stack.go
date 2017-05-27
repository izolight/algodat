package datastructures

import (
	"fmt"
)

// Stack has the a pointer to the top element in the Stack, the current size and the maximum allowed size
type Stack struct {
	top     *Element
	size    int
	maxSize int
}

// Element consists of the value and a pointer to the next Element in the stack
type Element struct {
	value int
	next  *Element
}

func (s *Stack) push(val int) error {
	if s.isFull() {
		return fmt.Errorf("Can't push to full stack. Size: %d", s.size)
	}
	top := Element{val, s.top}
	s.top = &top
	s.size++
	return nil
}

func (s *Stack) pop() (int, error) {
	if !s.isEmpty() {
		val := s.top.value
		s.top = s.top.next
		s.size--
		return val, nil
	}
	return 0, fmt.Errorf("Can't pop from empty stack")
}

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
