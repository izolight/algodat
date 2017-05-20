package main

import (
	"izolight/algorithms_datastructures/graph"
	"izolight/algorithms_datastructures/queue"
)

func breadthSearch(g *graph.Graph, q *queue.Queue) {
	g.vertices[]
}

func main() {
	S, A, B, C, D, E, F, G := 0, 1, 2,3, 4, 5, 6, 7

	myGraph := graph.Graph{{S, A, B, C, D, E, F, G},
		{
			{&S, &A},
			{&S, &B},
			{&S, &C},
			{&A, &D},
			{&B, &E},
			{&C, &F},
			{&D, &G},
			{&E, &G},
			{&F, &G},
		}}

	myQueue := queue.Queue{}
	breadthSearch(myGraph, myQueue)
}