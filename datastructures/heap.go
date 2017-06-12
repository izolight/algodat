package datastructures

import (
	"fmt"
)

type Heap struct {
	root   *HeapNode
	end    *HeapNode
	height int
	size   int
}

type MinHeap Heap
type MaxHeap Heap

type HeapNode struct {
	data       fmt.Stringer
	parent     *HeapNode
	leftChild  *HeapNode
	rightChild *HeapNode
}

func (h *Heap) Empty() bool {
	if h.root != nil {
		return false
	}
	return true
}

func (h *Heap) Size() int {
	// TODO
	return 0
}

func (h *Heap) Elements() []fmt.Stringer {
	// TODO
	return nil
}

func (n *HeapNode) isInternal() bool {
	// TODO
	return false
}

func (n *HeapNode) isExternal() bool {
	// TODO
	return false
}

func (n *HeapNode) isRoot() bool {
	// TODO
	return false
}

func (h *Heap) Insert(val fmt.Stringer) {
	if h.root == nil {
		node := HeapNode{val, nil, nil, nil}
		h.root = &node
		h.end = &node
		h.size++
		return
	}

}

func (h *MinHeap) Insert(val fmt.Stringer) {
	if h.root == nil {
		node := HeapNode{val, nil, nil, nil}
		h.root = &node
		h.end = &node
		return
	}

	if h.end.leftChild == nil {
		node := HeapNode{val, h.end, nil, nil}
	}
}

func (h *MaxHeap) Insert(val fmt.Stringer) {
	return
}
