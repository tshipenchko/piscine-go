package piscine

func ListLast(l *List) interface{} {
	node := l.Head

	if node == nil {
		return nil
	}

	for node.Next != nil {
		node = node.Next
	}

	return node.Data
}
