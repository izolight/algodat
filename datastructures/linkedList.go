package datastructures

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

var linkedListTemplates = template.Must(template.ParseFiles("templates/header.html", "templates/footer.html", "templates/linkedlist.html"))

// LinkedList represents the linked list datastructure
type LinkedList struct {
	head *Node
}

// Node is a representation of a node in a linked list
type Node struct {
	value int
	next  *Node
}

type linkedListData struct {
	Address, Next string
	Value         int
}

// NewLinkedList returns an empty LinkedList
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

func (l *LinkedList) search(val int) (*Node, error) {
	node := l.head
	for node != nil {
		if node.value == val {
			return node, nil
		}
		node = node.next
	}
	return nil, fmt.Errorf("Couldn't find %d in list", val)
}

func (l *LinkedList) delete(val int) error {
	if l.head == nil {
		return fmt.Errorf("%d not found in list", val)
	}
	if l.head.value == val {
		l.head = l.head.next
		return nil
	}
	node := l.head
	for node.next != nil {
		if node.next.value == val {
			node.next = node.next.next
			return nil
		}
		node = node.next
	}
	return fmt.Errorf("%d not found in list", val)
}

// View displays all values in the linked list
func (l *LinkedList) View(w http.ResponseWriter, r *http.Request) {
	var data []linkedListData
	node := l.head
	for node != nil {
		data = append(data, linkedListData{fmt.Sprintf("%p", node), fmt.Sprintf("%p", node.next), node.value})
		node = node.next
	}
	view := viewData{"Linked List", data, errors}
	err := linkedListTemplates.ExecuteTemplate(w, "linkedlist", view)
	if err != nil {
		fmt.Fprintf(w, "Could not render template: %v", err)
	}
	errors = errors[:0]
}

// InsertAtFront inserts the form value at the front of the linked list
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

// InsertAtEnd inserts the form value at the end of the linked list
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

// DeleteFromFront deletes the first node of the linked list
func (l *LinkedList) DeleteFromFront(w http.ResponseWriter, r *http.Request) {
	err := l.deleteFromFront()
	if err != nil {
		errors = append(errors, fmt.Sprintf("%v", err))
	}
	http.Redirect(w, r, "/linkedlist", http.StatusSeeOther)
}

// DeleteFromEnd deletes the last node of the linked list
func (l *LinkedList) DeleteFromEnd(w http.ResponseWriter, r *http.Request) {
	err := l.deleteFromEnd()
	if err != nil {
		errors = append(errors, fmt.Sprintf("%v", err))
	}
	http.Redirect(w, r, "/linkedlist", http.StatusSeeOther)
}

func (l *LinkedList) Search(w http.ResponseWriter, r *http.Request) {
	search, err := strconv.Atoi(r.FormValue("search"))
	if err != nil {
		errors = append(errors, fmt.Sprintf("Could not parse form value: %v", err))
	} else {
		node, err := l.search(search)
		if err != nil {
			errors = append(errors, fmt.Sprintf("%v", err))
		}
		if node != nil {
			errors = append(errors, fmt.Sprintf("%d found at address %p", search, node))
		}
	}
	http.Redirect(w, r, "/linkedlist", http.StatusSeeOther)
}

func (l *LinkedList) Delete(w http.ResponseWriter, r *http.Request) {
	delete, err := strconv.Atoi(r.FormValue("delete"))
	if err != nil {
		errors = append(errors, fmt.Sprintf("Could not parse form value: %v", err))
	} else {
		err = l.delete(delete)
		if err != nil {
			errors = append(errors, fmt.Sprintf("%v", err))
		}
	}
	http.Redirect(w, r, "/linkedlist", http.StatusSeeOther)
}
