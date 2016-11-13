package stack

import (
	"fmt"
)

type Stack struct {
	top  *Node
	size int
}

type Node struct {
	value int
	next  *Node
}

func Push(s *Stack, val int) {
	top := Node{val, s.top}
	s.top = &top
	s.size++
}

func Pop(s *Stack) {
	if !IsEmpty(s) {
		s.top = s.top.next
		s.size--
	}
}

func Top(s *Stack) int {
	return s.top.value
}

func Size(s *Stack) int {
	return s.size
}

func IsEmpty(s *Stack) bool {
	if s.size > 0 {
		return false
	}
	return true
}

func PrintStack(s *Stack) {
	if !IsEmpty(s) {
		fmt.Printf("Top: %d\tsize: %d\tEmpty: %t\n", Top(s), Size(s), IsEmpty(s))
	} else {
		fmt.Printf("Size: %d\tEmpty: %t\n", Size(s), IsEmpty(s))
	}
}