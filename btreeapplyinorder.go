package piscine

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if f == nil {
		return
	}

	if root == nil {
		return
	}

	BTreeApplyInorder(root.Left, f)

	_, _ = f(root.Data)

	BTreeApplyInorder(root.Right, f)
}
