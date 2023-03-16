package piscine

func BTreeApplyPostorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root.Left != nil {
		BTreeApplyInorder(root.Left, f)
	}

	if root.Right != nil {
		BTreeApplyInorder(root.Right, f)
	}

	f(root.Data)
}
