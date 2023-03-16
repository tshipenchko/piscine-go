package piscine

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return BTreeLevelCountHelper(root, 0)
}

func BTreeLevelCountHelper(root *TreeNode, count int) int {
	if root == nil {
		return count
	}

	leftCount := BTreeLevelCountHelper(root.Left, count+1)
	RightCount := BTreeLevelCountHelper(root.Right, count+1)

	if leftCount > RightCount {
		return leftCount
	}
	return RightCount
}
