package piscine

import "fmt"

func PrintTree(root *TreeNode) {
	if root == nil {
		return
	}

	var printNode func(node *TreeNode, prefix string, isLeft bool)
	printNode = func(node *TreeNode, prefix string, isLeft bool) {
		if node == nil {
			return
		}

		if node.Right != nil {
			printNode(node.Right, prefix+map[bool]string{true: "│ ", false: "  "}[isLeft], false)
		}

		fmt.Printf("%s%s%s\n", prefix+map[bool]string{true: "└─", false: "┌─"}[isLeft], node.Data, map[bool]string{true: "", false: "─┐"}[isLeft])

		if node.Left != nil {
			printNode(node.Left, prefix+map[bool]string{true: "  ", false: "│ "}[isLeft], true)
		}
	}

	printNode(root, "", true)
}
