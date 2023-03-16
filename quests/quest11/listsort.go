package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	size := 0
	node := l

	for node != nil {
		size++
		node = node.Next
	}

	for i := 0; i < size-1; i++ {
		curr := l
		for j := 0; j < size-i-1; j++ {
			if curr.Data > curr.Next.Data {
				curr.Data, curr.Next.Data = curr.Next.Data, curr.Data
			}
			curr = curr.Next
		}
	}

	return l
}
