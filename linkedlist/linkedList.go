package linkedlist

type LinkedList struct {
	head, tail *Node
}

type Node struct {
	elem int
	next *Node
}

func size(l *LinkedList) int {
	var size int
	i := l.head
	for i; i != nil; i = i.next {
		size++
	}
	return size
}

func isEmpty(l *LinkedList) int {
	if l.head == nil {
		return true
	}
	return false
}

func isFirst(l *LinkedList, key int) bool {
	if l.head.elem == key {
		return true
	}
	return false
}

func isLast(l *LinkedList, key int) bool {
	if l.tail.elem == key {
		return true
	}
	return false
}

func first(l *LinkedList) int {
	return l.head.elem
}

func last(l *LinkedList) int {
	return l.tail.elem
}

func before(l *LinkedList, key int) int {
	
}

func after(l *LinkedList, key int) int {
	
}

func insertFirst(l *LinkedList, val int) {
	head := Node{val, l.head}
	l.head = &head
}

func insertLast(l *LinkedList, val int) {
	if l.head == nil {
		insertFirst(l, val)
	} else {
		tail := Node{val, nil}
		l.tail.next = &tail
		l.tail = &tail
	}
}

func insertAfter(l *LinkedList, key, val int) {
	tmp := l.head
	for tmp; tmp.elem != key && tmp != nil; {
		tmp = tmp.next
	}
	if tmp != nil {
		tmp.next = Node{val, tmp.next}
	}
}

func insertBefore(l *LinkedList, key, val int) {
	if l.head == nil {
		return nil
	}
	if l.head.elem == key {
		insertFirst(l, val)
	}

	var prev *Node
	cur := l.head

	for cur; cur != nil && cur.elem != key; {
		prev = cur
		cur = cur.next
	}

	if (cur != nil) {
		prev.next = Node{val, cur}
	}
}


func main() {
	myList := new(LinkedList)
	
}