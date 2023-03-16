package piscine

import "fmt"

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	notNode := BTreeDeleteNodeHelper2(root, node)
	if notNode == nil {
		return nil
	}
	fmt.Println("sd", notNode.Data)

	out := &TreeNode{Data: notNode.Data}
	BTreeDeleteNodeHelper(root, node, root, out)

	root.Data = out.Data
	root.Left = out.Left
	root.Right = out.Right
	return out
}

func BTreeDeleteNodeHelper(root, node, node2, out *TreeNode) {
	if root == nil {
		return
	}

	if root != node || root != node2 {
		BTreeInsertData(out, root.Data)
	}

	BTreeDeleteNodeHelper(root.Left, node, node2, out)
	BTreeDeleteNodeHelper(root.Right, node, node2, out)
}

func BTreeDeleteNodeHelper2(root, node *TreeNode) *TreeNode {
	if root != node {
		return root
	}
	left := BTreeDeleteNodeHelper2(root.Left, node)
	if left != nil {
		return left
	}
	right := BTreeDeleteNodeHelper2(root.Right, node)
	if right != nil {
		return right
	}

	return nil
}
