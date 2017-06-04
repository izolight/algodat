package datastructures

import "testing"

func TestAddToQueue(t *testing.T) {
	q := NewQueue(5)
	var val = []IntTest{1, 2, 3}
	for _, i := range val {
		err := q.Enqueue(IntTest(i))
		if err != nil {
			t.Errorf("%v", err)
		}
		back, err := q.Back()
		if err != nil {
			t.Errorf("%v", err)
		}
		if back != IntTest(i) {
			t.Errorf("Expected back: %v, got %v", i, back)
		}
	}
	front, err := q.Front()
	if err != nil {
		t.Errorf("%v", err)
	}
	if front != val[0] {
		t.Errorf("Expected back: %v, got %v", val[0], front)
	}

	if q.Size() != len(val) {
		t.Errorf("Expected size: %v, got %v", len(val), q.Size())
	}
}

func TestRemoveFromQueue(t *testing.T) {
	q := NewQueue(5)
	var val = []IntTest{1, 2, 3}
	for _, i := range val {
		err := q.Enqueue(IntTest(i))
		if err != nil {
			t.Errorf("%v", err)
		}
	}
	dequeued, err := q.Dequeue()
	if err != nil {
		t.Errorf("%v", err)
	}
	if dequeued != val[0] {
		t.Errorf("Expected popped: %v, got %v", val[0], dequeued)
	}
	if q.Size() != len(val)-1 {
		t.Errorf("Expected queue size: %v, size %v", len(val)-1, q.Size())
	}
	q.Dequeue()
	q.Dequeue()
	if !q.Empty() {
		t.Errorf("Expected empty queue, size %v", q.Size())
	}
}

func TestEmptyQueue(t *testing.T) {
	q := NewQueue(0)
	front, err := q.Front()
	if err == nil {
		t.Errorf("Expected empty queue, got %v", front)
	}
	err = nil
	dequeued, err := q.Dequeue()
	if err == nil {
		t.Errorf("Expected empty queue, got %v", dequeued)
	}
	back, err := q.Back()
	if err == nil {
		t.Errorf("Expected empty queue, got %v", back)
	}
}

func TestFullQueue(t *testing.T) {
	q := NewQueue(1)
	var val IntTest = 1
	q.Enqueue(val)
	err := q.Enqueue(val)
	if err == nil {
		t.Errorf("Expected full stack, size %v", q.Size())
	}
}

func TestShowQueue(t *testing.T) {
	q := NewQueue(3)
	var val = []IntTest{1, 2, 3}
	for _, i := range val {
		err := q.Enqueue(IntTest(i))
		if err != nil {
			t.Errorf("%v", err)
		}
	}
	elements := q.Elements()
	for i := range val {
		if elements[i] != val[i] {
			t.Errorf("Expected %v, got %v", val[i], elements[i])
		}
	}
}
