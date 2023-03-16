package piscine

func BTreeApplyPreorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}

	f(root.Data)
	BTreeApplyPreorder(root.Left, f)
	BTreeApplyPostorder(root.Right, f)
}
