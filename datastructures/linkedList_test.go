package datastructures

import "testing"

func TestAddToLinkedList(t *testing.T) {
	l := NewLinkedList()
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

func TestRemoveFromLinkedList(t *testing.T) {
	l := NewLinkedList()
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

func TestEmptyLinkedList(t *testing.T) {
	l := NewLinkedList()
	head, err := l.Head()
	if err == nil {
		t.Errorf("Expected empty queue, got %v", head)
	}
	tail, err := l.Tail()
	if err == nil {
		t.Errorf("Expected empty queue, got %v", tail)
	}
	removedHead, err := l.DeleteFromFront()
	if err == nil {
		t.Errorf("Expected empty queue, got %v", removedHead)
	}
	removedTail, err := l.DeleteFromEnd()
	if err == nil {
		t.Errorf("Expected empty queue, got %v", removedTail)
	}
}
