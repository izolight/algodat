package datastructures

// DoublyLinkedList represents the doubly linked list datastructure
type DoublyLinkedList struct {
	head     *DoublyLinkedListNode
	messages []string
}

// DoublyLinkedListNode represents a node in a doubly linked list
type DoublyLinkedListNode struct {
	value          int
	previous, next *DoublyLinkedListNode
}

// NewDoublyLinkedList returns an empty DoublyLinkedList
func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{nil, nil}
}

func (l *DoublyLinkedList) insertAtFront(val int) {

}

func (l *DoublyLinkedList) insertAtEnd(val int) {

}

func (l *DoublyLinkedList) deleteFromFront() {

}

func (l *DoublyLinkedList) deleteFromEnd() {

}

func (l *DoublyLinkedList) insertAfter(val int) {

}

func (l *DoublyLinkedList) delete(val int) {

}
