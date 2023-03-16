package piscine

func BTreeIsBinary(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left == nil {
		return true
	}
	if !(root.Data > root.Left.Data) {
		return false
	}
	if root.Right == nil {
		return true
	}
	if !(root.Data < root.Right.Data) {
		return false
	}

	return BTreeIsBinary(root.Left) && BTreeIsBinary(root.Right)
}
