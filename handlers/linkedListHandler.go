type linkedListData struct {
	Address, Next string
	Value         int
}

var linkedListTemplates = template.Must(template.ParseFiles("templates/header.html", "templates/footer.html", "templates/linkedlist.html"))

// View displays all values in the linked list
func (l *LinkedList) View(w http.ResponseWriter, r *http.Request) {
	var data []linkedListData
	node := l.head
	for node != nil {
		data = append(data, linkedListData{fmt.Sprintf("%p", node), fmt.Sprintf("%p", node.next), node.value})
		node = node.next
	}
	view := viewData{"Linked List", data, l.messages, l.errors}
	err := linkedListTemplates.ExecuteTemplate(w, "linkedlist", view)
	if err != nil {
		fmt.Fprintf(w, "Could not render template: %v", err)
	}
	l.messages = l.messages[:0]
	l.errors = l.errors[:0]
}

// InsertAtFront inserts the form value at the front of the linked list
func (l *LinkedList) InsertAtFront(w http.ResponseWriter, r *http.Request) {
	new, err := strconv.Atoi(r.FormValue("new"))
	if err != nil {
		l.errors = append(l.errors, fmt.Sprintf("Could not parse form value: %v", err))
	} else {
		err = l.insertAtFront(new)
		if err != nil {
			l.errors = append(l.errors, fmt.Sprintf("%v", err))
		} else {
			l.messages = append(l.messages, fmt.Sprintf("Inserted %d at front", new))
		}
	}
	http.Redirect(w, r, "/linkedlist", http.StatusSeeOther)
}

// InsertAtEnd inserts the form value at the end of the linked list
func (l *LinkedList) InsertAtEnd(w http.ResponseWriter, r *http.Request) {
	new, err := strconv.Atoi(r.FormValue("new"))
	if err != nil {
		l.errors = append(l.errors, fmt.Sprintf("Could not parse form value: %v", err))
	} else {
		err = l.insertAtEnd(new)
		if err != nil {
			l.errors = append(l.errors, fmt.Sprintf("%v", err))
		} else {
			l.messages = append(l.messages, fmt.Sprintf("Inserted %d at end", new))
		}
	}
	http.Redirect(w, r, "/linkedlist", http.StatusSeeOther)
}

// DeleteFromFront deletes the first node of the linked list
func (l *LinkedList) DeleteFromFront(w http.ResponseWriter, r *http.Request) {
	deleted, err := l.deleteFromFront()
	if err != nil {
		l.errors = append(l.errors, fmt.Sprintf("%v", err))
	} else {
		l.messages = append(l.messages, fmt.Sprintf("Deleted %d with address %p at front", deleted.value, deleted))
	}
	http.Redirect(w, r, "/linkedlist", http.StatusSeeOther)
}

// DeleteFromEnd deletes the last node of the linked list
func (l *LinkedList) DeleteFromEnd(w http.ResponseWriter, r *http.Request) {
	deleted, err := l.deleteFromEnd()
	if err != nil {
		l.errors = append(l.errors, fmt.Sprintf("%v", err))
	} else {
		l.messages = append(l.messages, fmt.Sprintf("Deleted %d at wit address %p at end", deleted.value, deleted))
	}
	http.Redirect(w, r, "/linkedlist", http.StatusSeeOther)
}

// Search searches the form value and returns the address of it
func (l *LinkedList) Search(w http.ResponseWriter, r *http.Request) {
	search, err := strconv.Atoi(r.FormValue("search"))
	if err != nil {
		l.errors = append(l.errors, fmt.Sprintf("Could not parse form value: %v", err))
	} else {
		node, err := l.search(search)
		if err != nil {
			l.errors = append(l.errors, fmt.Sprintf("%v", err))
		} else {
			l.messages = append(l.messages, fmt.Sprintf("%d found at address %p", search, node))
		}
	}
	http.Redirect(w, r, "/linkedlist", http.StatusSeeOther)
}

// Delete searches for the form value and deletes it if found
func (l *LinkedList) Delete(w http.ResponseWriter, r *http.Request) {
	delete, err := strconv.Atoi(r.FormValue("delete"))
	if err != nil {
		l.errors = append(l.errors, fmt.Sprintf("Could not parse form value: %v", err))
	} else {
		deleted, err := l.delete(delete)
		if err != nil {
			l.errors = append(l.errors, fmt.Sprintf("%v", err))
		} else {
			l.messages = append(l.messages, fmt.Sprintf("Deleted %d at address %p", deleted.value, deleted))
		}
	}
	http.Redirect(w, r, "/linkedlist", http.StatusSeeOther)
}