package piscine

import "fmt"

func Show(l *List) {
	node := l.Head

	for node != nil {
		fmt.Print(node.Data)
		fmt.Print(" -> ")
		node = node.Next
	}
	fmt.Print(node)
	fmt.Println()
}
