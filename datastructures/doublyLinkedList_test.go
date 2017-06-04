package datastructures

import (
	"testing"
)

func TestAddToDoublyLinkedList(t *testing.T) {
	l := NewDoublyLinkedList()
	var val = []IntTest{1, 2, 3}
	for _, i := range val {
		err := l.InsertAtEnd(IntTest(i))
		if err != nil {
			t.Errorf("%v", err)
		}
		tail, err := l.Tail()
		if err != nil {
			t.Errorf("%v", err)
		}
		if tail != IntTest(i) {
			t.Errorf("Expected tail: %v, got %v", i, tail)
		}
	}
	head, err := l.Head()
	if err != nil {
		t.Errorf("%v", err)
	}
	if head != val[0] {
		t.Errorf("Expected back: %v, got %v", val[0], head)
	}

	if l.Size() != len(val) {
		t.Errorf("Expected size: %v, got %v", len(val), l.Size())
	}
}

func TestRemoveFromDoublyLinkedList(t *testing.T) {
	l := NewDoublyLinkedList()
	var val = []IntTest{1, 2, 3}
	for _, i := range val {
		err := l.InsertAtFront(IntTest(i))
		if err != nil {
			t.Errorf("%v", err)
		}
	}
	tail, err := l.DeleteFromEnd()
	if err != nil {
		t.Errorf("%v", err)
	}
	if tail.data != val[0] {
		t.Errorf("Expected tail: %v, got %v", val[len(val)-1], tail.data)
	}
	head, err := l.DeleteFromFront()
	if err != nil {
		t.Errorf("%v", err)
	}
	if head.data != val[len(val)-1] {
		t.Errorf("Expected head: %v, got %v", val[0], head.data)
	}

	if l.Size() != len(val)-2 {
		t.Errorf("Expected list size: %v, size %v", len(val)-2, l.Size())
	}
	l.DeleteFromFront()
	if !l.Empty() {
		t.Errorf("Expected empty list, size %v", l.Size())
	}
	l.InsertAtFront(IntTest(1))
	l.DeleteFromEnd()
	if !l.Empty() {
		t.Errorf("Expected empty list, size %v", l.Size())
	}
}

func TestEmptyDoublyLinkedList(t *testing.T) {
	l := NewDoublyLinkedList()
	head, err := l.Head()
	if err == nil {
		t.Errorf("Expected empty list, got %v", head)
	}
	tail, err := l.Tail()
	if err == nil {
		t.Errorf("Expected empty list, got %v", tail)
	}
	removedHead, err := l.DeleteFromFront()
	if err == nil {
		t.Errorf("Expected empty list, got %v", removedHead)
	}
	removedTail, err := l.DeleteFromEnd()
	if err == nil {
		t.Errorf("Expected empty list, got %v", removedTail)
	}
}

func TestSearchAndDeleteDoublyLinkedList(t *testing.T) {
	l := NewDoublyLinkedList()
	var val = []IntTest{1, 2, 3, 4, 5}
	for i := range val {
		err := l.InsertAtFront(IntTest(val[i]))
		if err != nil {
			t.Errorf("%v", err)
		}
	}
	search, err := l.Search(val[1])
	if err != nil {
		t.Errorf("%v", err)
	}
	if search.data != val[1] {
		t.Errorf("Expected %v, got %v", val[1], search.data)
	}
	deleted, err := l.Delete(val[2])
	if err != nil {
		t.Errorf("%v", err)
	}
	if deleted.data != val[2] {
		t.Errorf("Expected %v, got %v", val[1], deleted.data)
	}
	if l.Size() != len(val)-1 {
		t.Errorf("Expected size %v, got %v", len(val)-1, l.Size())
	}
	// this is the front element so it triggers a different codepath
	deleted, err = l.Delete(val[len(val)-1])
	if err != nil {
		t.Errorf("%v", err)
	}
	if deleted.data != val[len(val)-1] {
		t.Errorf("Expected %v, got %v", val[len(val)-1], deleted.data)
	}
	// try to delete deleted element
	deleted, err = l.Delete(val[2])
	if err == nil {
		t.Errorf("Expected not found, got %v", deleted)
	}
	search, err = l.Search(val[2])
	if err == nil {
		t.Errorf("Expected not found, got %v", search)
	}
}
