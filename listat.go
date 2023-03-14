package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	index := 0
	node := l
	for node != nil {
		if index == pos {
			return node
		}
		node = node.Next
		index++
	}

	return nil
}
