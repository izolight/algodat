package datastructures

import (
	"fmt"
)

// LinkedList represents the linked list datastructure
type LinkedList struct {
	head *Node
	tail *Node
	size int
}

// Node is a representation of a node in a linked list
type Node struct {
	data fmt.Stringer
	next *Node
}

// Empty checks whether the linked list is empty
func (l *LinkedList) Empty() bool {
	if l.head != nil {
		return false
	}
	return true
}

func (l *LinkedList) Size() int {
	return l.size
}

func (l *LinkedList) Elements() []fmt.Stringer {
	// TODO
	return nil
}

// Head returns the first element in the list
func (l *LinkedList) Head() (fmt.Stringer, error) {
	if l.Empty() {
		return nil, fmt.Errorf("List is empty")
	}
	return l.head.data, nil
}

// Tail returns the last element in the list
func (l *LinkedList) Tail() (fmt.Stringer, error) {
	if l.Empty() {
		return nil, fmt.Errorf("List is empty")
	}
	return l.tail.data, nil
}

// NewLinkedList returns an empty LinkedList
func NewLinkedList() *LinkedList {
	return &LinkedList{nil, nil, 0}
}

// InsertAtFront adds an element before the head
func (l *LinkedList) InsertAtFront(val fmt.Stringer) error {
	node := Node{val, l.head}
	if l.Empty() {
		l.tail = &node
	} else if l.size == 1 {
		l.tail = l.head
	}
	l.head = &node
	return nil
}

// InsertAtEnd adds an element after the tail
func (l *LinkedList) insertAtEnd(val fmt.Stringer) error {
	node := Node{val, nil}
	if l.head == nil {
		return l.InsertAtFront(val)
	}
	previous, current := l.head, l.head
	for current != nil {
		previous = current
		current = current.next
	}
	previous.next = &node
	l.tail = &node
	return nil
}

func (l *LinkedList) deleteFromFront() (*Node, error) {
	if l.head == nil {
		return nil, fmt.Errorf("Can't delete from empty list")
	}
	deleted := l.head
	l.head = l.head.next
	return deleted, nil
}

func (l *LinkedList) deleteFromEnd() (*Node, error) {
	if l.head == nil {
		return nil, fmt.Errorf("Can't delete from empty list")
	}
	if l.head.next == nil {
		deleted := l.head
		l.head = nil
		return deleted, nil
	}
	previous := l.head
	for previous.next.next != nil {
		previous = previous.next
	}
	deleted := previous.next
	previous.next = nil
	return deleted, nil
}

func (l *LinkedList) search(val fmt.Stringer) (*Node, error) {
	node := l.head
	for node != nil {
		if node.data == val {
			return node, nil
		}
		node = node.next
	}
	return nil, fmt.Errorf("Can't find %d in list", val)
}

func (l *LinkedList) delete(val fmt.Stringer) (*Node, error) {
	if l.head == nil {
		return nil, fmt.Errorf("Can't find %d in list", val)
	}
	if l.head.data == val {
		deleted := l.head
		l.head = l.head.next
		return deleted, nil
	}
	node := l.head
	for node.next != nil {
		if node.next.data == val {
			deleted := node.next
			node.next = node.next.next
			return deleted, nil
		}
		node = node.next
	}
	return nil, fmt.Errorf("Can't find %d in list", val)
}
