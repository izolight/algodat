package datastructures

import (
	"fmt"
)

// DoublyLinkedList represents the doubly linked list datastructure
type DoublyLinkedList struct {
	head *DoublyLinkedListNode
	tail *DoublyLinkedListNode
	size int
}

// DoublyLinkedListNode represents a node in a doubly linked list
type DoublyLinkedListNode struct {
	data     fmt.Stringer
	previous *DoublyLinkedListNode
	next     *DoublyLinkedListNode
}

// Empty checks whether the linked list is empty
func (l *DoublyLinkedList) Empty() bool {
	if l.head != nil {
		return false
	}
	return true
}

// Size returns the current size of the linked list
func (l *DoublyLinkedList) Size() int {
	return l.size
}

// Head returns the first element in the list
func (l *DoublyLinkedList) Head() (fmt.Stringer, error) {
	if l.Empty() {
		return nil, fmt.Errorf("List is empty")
	}
	return l.head.data, nil
}

// Tail returns the last element in the list
func (l *DoublyLinkedList) Tail() (fmt.Stringer, error) {
	if l.Empty() {
		return nil, fmt.Errorf("List is empty")
	}
	return l.tail.data, nil
}

// NewDoublyLinkedList returns an empty DoublyLinkedList
func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{nil, nil, 0}
}

// InsertAtFront adds an element before the head
func (l *DoublyLinkedList) InsertAtFront(val fmt.Stringer) error {
	node := DoublyLinkedListNode{val, nil, l.head}
	if l.Empty() {
		l.tail = &node
	} else if l.size == 1 {
		l.tail = l.head
	}
	if !l.Empty() {
		l.head.previous = &node
	}
	l.head = &node
	l.size++
	return nil
}

// InsertAtEnd adds an element after the tail
func (l *DoublyLinkedList) InsertAtEnd(val fmt.Stringer) error {
	if l.head == nil {
		return l.InsertAtFront(val)
	}
	node := DoublyLinkedListNode{val, l.tail, nil}
	l.tail.next = &node
	l.tail = &node
	l.size++
	return nil
}

// DeleteFromFront removes and returns the first element
func (l *DoublyLinkedList) DeleteFromFront() (*DoublyLinkedListNode, error) {
	if l.head == nil {
		return nil, fmt.Errorf("Can't delete from empty list")
	}
	deleted := l.head
	l.size--
	if l.Size() == 1 {
		l.head = nil
		l.tail = nil
		return deleted, nil
	}
	l.head = l.head.next
	l.head.previous = nil
	return deleted, nil
}

// DeleteFromEnd removes and returns the last element
func (l *DoublyLinkedList) DeleteFromEnd() (*DoublyLinkedListNode, error) {
	if l.head == nil {
		return nil, fmt.Errorf("Can't delete from empty list")
	}
	if l.tail.previous == nil {
		deleted := l.tail
		l.head = nil
		l.tail = nil
		l.size--
		return deleted, nil
	}
	deleted := l.tail
	l.tail = l.tail.previous
	l.tail.next = nil
	l.size--
	return deleted, nil
}

// InsertAfter adds an element after the found value
func (l *DoublyLinkedList) InsertAfter(val fmt.Stringer, search fmt.Stringer) error {
	return nil
}

// Delete searches for a value in the list and deletes the first it finds
func (l *DoublyLinkedList) Delete(val fmt.Stringer) (*DoublyLinkedListNode, error) {
	if l.head == nil {
		return nil, fmt.Errorf("Can't find %d in list", val)
	}
	if l.head.data == val {
		return l.DeleteFromFront()
	}
	if l.tail.data == val {
		return l.DeleteFromEnd()
	}
	node := l.head
	for node != nil {
		if node.data == val {
			deleted := node
			node.previous.next = node.next
			l.size--
			return deleted, nil
		}
		node = node.next
	}
	return nil, fmt.Errorf("Can't find %d in list", val)
}

// Search searches for a value in the list and returns the first it finds
func (l *DoublyLinkedList) Search(val fmt.Stringer) (*DoublyLinkedListNode, error) {
	if l.head.data == val {
		return l.head, nil
	}
	if l.tail.data == val {
		return l.tail, nil
	}
	node := l.head
	for node != nil {
		if node.data == val {
			return node, nil
		}
		node = node.next
	}
	return nil, fmt.Errorf("Can't find %d in list", val)
}
