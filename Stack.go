package main

import (
	"fmt"
	"time"
)

type Stack struct {
	top  *Node
	size int
}

type Node struct {
	value int
	next  *Node
}

func push(s *Stack, val int) {
	top := Node{val, s.top}
	s.top = &top
	s.size++
}

func pop(s *Stack) {
	if !isEmpty(s) {
		s.top = s.top.next
		s.size--
	}
}

func top(s *Stack) int {
	return s.top.value
}

func size(s *Stack) int {
	return s.size
}

func isEmpty(s *Stack) bool {
	if s.size > 0 {
		return false
	}
	return true
}

func printStack(s *Stack) {
	if !isEmpty(s) {
		fmt.Printf("Top: %d\tsize: %d\tEmpty: %t\n", top(s), size(s), isEmpty(s))
	} else {
		fmt.Printf("Size: %d\tEmpty: %t\n", size(s), isEmpty(s))
	}
}

func main() {
	myStack := new(Stack)
	const max int = 16 * 1024 * 1024 + 2
	start := time.Now()
	for i := 0; i < max; i++ {
		push(myStack, i)
	}
	end := time.Now()
	printStack(myStack)
	fmt.Printf("Time taken to store %d elements: %v\n", max, end.Sub(start))
	start = time.Now()
	for i := 0; i < max / 2; i++ {
		pop(myStack)
	}
	end = time.Now()
	printStack(myStack)
	fmt.Printf("Time taken to remove %d elements: %v", max / 2, end.Sub(start))
}