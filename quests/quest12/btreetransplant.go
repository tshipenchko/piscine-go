package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if root == nil || node == nil {
		return root
	}
	node.Data = rplc.Data
	node.Left = rplc.Left
	node.Right = rplc.Right
	return root
}
