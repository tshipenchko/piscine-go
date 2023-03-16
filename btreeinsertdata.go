package piscine

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		root = &TreeNode{Data: data, Parent: root}
		return root
	}

	if root.Data <= data { // Node is less than value. So write it right
		if root.Right != nil {
			BTreeInsertData(root.Right, data)
		}

		root.Right = &TreeNode{Data: data, Parent: root}
	} else { // root.Data < data
		if root.Left != nil {
			BTreeInsertData(root.Left, data)
		}

		root.Left = &TreeNode{Data: data, Parent: root}
	}

	return root
}
