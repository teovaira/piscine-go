package piscine

func BTreeApplyPreorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if f == nil {
		return
	}

	if root == nil {
		return
	}

	_, _ = f(root.Data)

	BTreeApplyPreorder(root.Left, f)

	BTreeApplyPreorder(root.Right, f)
}
