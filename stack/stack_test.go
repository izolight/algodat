package stack

import "testing"

func TestPush(t *testing.T) {
	s := new(Stack)
	max := 5
	for i := 0; i < max; i++ {
		Push(s, i)
	}

	size := Size(s)
	if size != max {
		t.Errorf("Expected %d, got %d", size, max)
	}

	top := Top(s)
	if top != max - 1 {
		t.Errorf("Top is %d, expected %d", top, max - 1)
	}

	empty := IsEmpty(s)
	if empty {
		t.Errorf("Expected %t, got %t", true, empty)
	}
}

func TestPop(t *testing.T) {
	s := new(Stack)
	max := 5
	for i := 0; i < max; i++ {
		Push(s, i)
	}
	toPop := max - 2

	for i := 0; i < toPop; i++ {
		Pop(s)
	}

	size := Size(s)
	if size != max - toPop {
		t.Errorf("Expected %d, got %d", size, max - toPop)
	}

	top := Top(s)
	if top != max - toPop - 1 {
		t.Errorf("Top is %d, expected %d", top, max - toPop - 1)
	}

	empty := IsEmpty(s)
	if empty {
		t.Errorf("Expected %t, got %t", true, empty)
	}
}