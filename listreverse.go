package piscine

func ListReverse(l *List) {
	// BRUH
	oldTail := l.Tail

	for l.Head != oldTail {
		Show(l)
		next := l.Head.Next

		// Move current Head to Tail
		l.Tail.Next = l.Head
		l.Tail = l.Head
		l.Tail.Next = nil

		// Set next as Head
		l.Head = next
	}
	Show(l)
}
