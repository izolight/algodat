package datastructures

type LinkedList struct {
	head *Node
}

type Node struct {
	value int
	next  *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{nil}
}

func (l *LinkedList) insertAtFront(val int) error {
	node := Node{val, l.head}
	l.head = &node
	return nil
}

func (l *LinkedList) insertAtEnd(val int) error {
	node := Node{val, nil}
	if l.head == nil {
		return l.insertAtFront(val)
	}
	previous, current := l.head, l.head
	for current != nil {
		previous = current
		current = current.next
	}
	previous.next = &node
	return nil
}

func (l *LinkedList) deleteFromFront() error {
	l.head = l.head.next
	return nil
}

func (l *LinkedList) deleteFromEnd() error {
	previous, current := l.head, l.head
	for current.next != nil {
		previous = current
		current = current.next
	}
	previous.next = nil
	return nil
}

func (l *LinkedList) search(val int) (bool, error) {
	return false, nil
}

func (l *LinkedList) delete(val int) error {
	return nil
}
