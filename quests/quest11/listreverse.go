package piscine

func ListReverse(l *List) {
	var prev, curr, next *NodeL

	curr = l.Head

	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	l.Head = prev
}
