package piscine

func ListForEachIf(l *List, f func(*NodeL), cond func(*NodeL) bool) {
	curr := l.Head

	for curr != nil {
		if cond(curr) {
			f(curr)
		}
		curr = curr.Next
	}
}

func IsPositiveNode(node *NodeL) bool {
	switch node.Data.(type) {
	case int, float32, float64, byte:
		return node.Data.(int) > 0
	default:
		return false
	}
}

func IsAlNode(node *NodeL) bool {
	switch node.Data.(type) {
	case int, float32, float64, byte:
		return false
	default:
		return true
	}
}
