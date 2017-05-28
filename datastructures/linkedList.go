package datastructures

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

var linkedListTemplates = template.Must(template.ParseFiles("templates/header.html", "templates/footer.html", "templates/linkedlist.html"))

type LinkedList struct {
	head *Node
}

type Node struct {
	value int
	next  *Node
}

type LinkedListData struct {
	Address, Next string
	Value         int
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
	if l.head == nil {
		return fmt.Errorf("Can't delete from empty list")
	}
	l.head = l.head.next
	return nil
}

func (l *LinkedList) deleteFromEnd() error {
	if l.head == nil {
		return fmt.Errorf("Can't delete from empty list")
	}
	if l.head.next == nil {
		l.head = nil
		return nil
	}
	previous := l.head
	for previous.next.next != nil {
		previous = previous.next
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

func (l *LinkedList) View(w http.ResponseWriter, r *http.Request) {
	var data []LinkedListData
	node := l.head
	for node != nil {
		data = append(data, LinkedListData{fmt.Sprintf("%p", node), fmt.Sprintf("%p", node.next), node.value})
		node = node.next
	}
	view := viewData{"Linked List", data, errors}
	err := linkedListTemplates.ExecuteTemplate(w, "linkedlist", view)
	if err != nil {
		fmt.Fprintf(w, "Could not render template: %v", err)
	}
	errors = errors[:0]
}

func (l *LinkedList) InsertAtFront(w http.ResponseWriter, r *http.Request) {
	new, err := strconv.Atoi(r.FormValue("new"))
	if err != nil {
		errors = append(errors, fmt.Sprintf("Could not parse form value: %v", err))
	} else {
		err = l.insertAtFront(new)
		if err != nil {
			errors = append(errors, fmt.Sprintf("%v", err))
		}
	}
	http.Redirect(w, r, "/linkedlist", http.StatusSeeOther)
}

func (l *LinkedList) InsertAtEnd(w http.ResponseWriter, r *http.Request) {
	new, err := strconv.Atoi(r.FormValue("new"))
	if err != nil {
		errors = append(errors, fmt.Sprintf("Could not parse form value: %v", err))
	} else {
		err = l.insertAtEnd(new)
		if err != nil {
			errors = append(errors, fmt.Sprintf("%v", err))
		}
	}
	http.Redirect(w, r, "/linkedlist", http.StatusSeeOther)
}

func (l *LinkedList) DeleteFromFront(w http.ResponseWriter, r *http.Request) {
	err := l.deleteFromFront()
	if err != nil {
		errors = append(errors, fmt.Sprintf("%v", err))
	}
	http.Redirect(w, r, "/linkedlist", http.StatusSeeOther)
}

func (l *LinkedList) DeleteFromEnd(w http.ResponseWriter, r *http.Request) {
	err := l.deleteFromEnd()
	if err != nil {
		errors = append(errors, fmt.Sprintf("%v", err))
	}
	http.Redirect(w, r, "/linkedlist", http.StatusSeeOther)
}
