package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	curr := root

	for curr != nil {
		if elem == curr.Data {
			return curr
		}
		if elem < curr.Data {
			curr = curr.Left
		} else {
			curr = curr.Right
		}
	}

	return nil
}
