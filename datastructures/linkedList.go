package datastructures

type LinkedList struct {
	head *Node
}

type Node struct {
	value int
	next  *Node
}

func (l *LinkedList) insert(val int) error {
	return nil
}

func (l *LinkedList) deleteFromFront() error {
	return nil
}

func (l *LinkedList) search(val int) (bool, error) {
	return false, nil
}

func (l *LinkedList) delete(val int) error {
	return nil
}
