package piscine

func ListSize(l *List) int {
	out := 0

	node := l.Head

	for node != nil {
		out++
		node = node.Next
	}

	return out
}
