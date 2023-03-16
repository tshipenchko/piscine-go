package piscine

func ListPushFront(l *List, data interface{}) {
	node := &NodeL{data, nil}

	if l.Head == nil && l.Tail == nil {
		l.Head = node
		l.Tail = node
	} else {
		node.Next = l.Head
		l.Head = node
	}
}
