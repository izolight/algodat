package main

import (
	"time"
	"fmt"
	"izolight/algorithms_datastructures/queue"
)

func main() {
	myQueue := new(queue.Queue)
	const max int = 16 * 1024 * 1024 + 2
	start := time.Now()
	for i := 0; i < max; i++ {
		queue.Enqueue(myQueue, i)
	}
	end := time.Now()
	fmt.Printf("Time taken to store %d elements: %v\n", max, end.Sub(start))
	queue.PrintQueue(myQueue)
	start = time.Now()
	for i := 0; i < max / 2; i++ {
		queue.Dequeue(myQueue)
	}
	end = time.Now()
	fmt.Printf("Time taken to remove %d elements: %v\n", max / 2, end.Sub(start))
	queue.PrintQueue(myQueue)
}
