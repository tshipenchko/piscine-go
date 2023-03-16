package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	levelCount := BTreeLevelCount(root)
	for i := 1; i <= levelCount; i++ {
		BTreeApplyByLevelHelper(root, f, i)
	}
}

func BTreeApplyByLevelHelper(root *TreeNode, f func(...interface{}) (int, error), level int) {
	if root == nil {
		return
	}
	if level == 1 {
		f(root.Data)
	} else {
		BTreeApplyByLevelHelper(root.Left, f, level-1)
		BTreeApplyByLevelHelper(root.Right, f, level-1)
	}
}
