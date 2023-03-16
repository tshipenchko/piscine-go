package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	newN := &NodeI{Data: data_ref, Next: nil}

	if l == nil || l.Data >= data_ref {
		newN.Next = l
		return newN
	}

	curr := l
	for curr.Next != nil && curr.Next.Data < data_ref {
		curr = curr.Next
	}

	newN.Next = curr.Next
	curr.Next = newN

	return l
}
