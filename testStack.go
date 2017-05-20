package main

import (
	"time"
	"fmt"
	"izolight/algorithms_datastructures/stack"
)

func main() {
	fmt.Println("Start Functional Test of stack")
	fmt.Println("Start Performance Test of stack")
	myStack := new(stack.Stack)
	const max int = 16 * 1024 * 1024 + 2
	start := time.Now()
	for i := 0; i < max; i++ {
		stack.Push(myStack, i)
	}
	end := time.Now()
	stack.PrintStack(myStack)
	fmt.Printf("Time taken to store %d elements: %v\n", max, end.Sub(start))
	start = time.Now()
	for i := 0; i < max / 2; i++ {
		stack.Pop(myStack)
	}
	end = time.Now()
	stack.PrintStack(myStack)
	fmt.Printf("Time taken to remove %d elements: %v", max / 2, end.Sub(start))
}