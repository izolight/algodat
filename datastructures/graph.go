package datastructures

type Graph struct {
	vertices []int
	edges []Edge
}

type Edge struct {
	v1, v2 *int;
}