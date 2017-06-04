package datastructures

import (
	"strconv"
	"testing"
)

type IntTest int

func (i IntTest) String() string {
	return strconv.Itoa(int(i))
}

func TestAddToStack(t *testing.T) {
	s := NewStack(5)
	var val = []IntTest{1, 2, 3}
	for _, i := range val {
		err := s.Push(IntTest(i))
		if err != nil {
			t.Errorf("%v", err)
		}
		top, err := s.Top()
		if err != nil {
			t.Errorf("%v", err)
		}
		if top != IntTest(i) {
			t.Errorf("Expected top: %v, got %v", i, top)
		}
	}

	if s.Size() != len(val) {
		t.Errorf("Expected size: %v, got %v", len(val), s.Size())
	}
}

func TestRemoveFromStack(t *testing.T) {
	s := NewStack(5)
	var val = []IntTest{1, 2, 3}
	for _, i := range val {
		err := s.Push(IntTest(i))
		if err != nil {
			t.Errorf("%v", err)
		}
	}
	popped, err := s.Pop()
	if err != nil {
		t.Errorf("%v", err)
	}
	if popped != val[len(val)-1] {
		t.Errorf("Expected popped: %v, got %v", val[len(val)-1], popped)
	}
	if s.Size() != len(val)-1 {
		t.Errorf("Expected stack size: %v, size %v", len(val)-1, s.Size())
	}
	s.Pop()
	s.Pop()
	if !s.Empty() {
		t.Errorf("Expected empty stack: %v, size %v", 1, s.Size())
	}
}

func TestEmptyStack(t *testing.T) {
	s := NewStack(0)
	top, err := s.Top()
	if err == nil {
		t.Errorf("Expected empty stack, got %v", top)
	}
	err = nil
	popped, err := s.Pop()
	if err == nil {
		t.Errorf("Expected empty stack, got %v", popped)
	}
}

func TestFullStack(t *testing.T) {
	s := NewStack(1)
	var val IntTest = 1
	s.Push(val)
	err := s.Push(val)
	if err == nil {
		t.Errorf("Expected full stack, size %v", s.Size())
	}
}

func TestShowStack(t *testing.T) {
	s := NewStack(3)
	var val = []IntTest{1, 2, 3}
	for _, i := range val {
		err := s.Push(IntTest(i))
		if err != nil {
			t.Errorf("%v", err)
		}
	}
	elements := s.Elements()
	for i := range val {
		if elements[i] != val[i] {
			t.Errorf("Expected %v, got %v", val[i], elements[i])
		}
	}
}
