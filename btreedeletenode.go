package piscine

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if node.Data < root.Data {
		root.Left = BTreeDeleteNode(root.Left, node)
	} else if node.Data > root.Data {
		root.Right = BTreeDeleteNode(root.Right, node)
	} else { // node.Data == root.Data
		if root.Left != nil && root.Right != nil {
			min := BTreeMin(root.Right)
			root.Data = min.Data
			root.Right = BTreeDeleteNode(root.Right, min)
		} else {
			if root.Left == nil && root.Right == nil {
				root = nil // Just delete
			} else if root.Right != nil {
				root = root.Right
			} else { // root.Left != nil
				root = root.Left
			}
		}
	}

	return root
}
