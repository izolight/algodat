package datastructures

import (
	"fmt"
)

type Element interface {
	fmt.Stringer
	Compare(Element) (int, error)
}

type Tree struct {
	data  interface{}
	left  *Tree
	right *Tree
}

func Insert(t *Tree, val interface{}) *Tree {
	if t == nil {
		return &Tree{val, nil, nil}
	}
	if val < t.data {
		t.left = Insert(t.left, val)
		return t
	}
}
