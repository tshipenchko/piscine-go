package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil || root.Data == elem {
		return root
	}

	if elem > root.Data {
		return BTreeSearchItem(root.Right, elem)
	}

	if elem < root.Data {
		return BTreeSearchItem(root.Left, elem)
	}

	return root
}
