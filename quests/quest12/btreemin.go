package piscine

func BTreeMin(root *TreeNode) *TreeNode {
	prev := root
	for root != nil {
		prev = root
		root = root.Left
	}
	return prev
}
