package piscine

func ListPushBack(l *List, data interface{}) {
	node := &NodeL{data, nil}

	if l.Head == nil && l.Tail == nil {
		l.Head = node
		l.Tail = node
	} else {
		l.Tail.Next = node
		l.Tail = node
	}
}
