package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	newNode := &TreeNode{Data: data}

	if root == nil {
		return newNode
	}

	curr := root
	for {
		if data < curr.Data {
			if curr.Left == nil {
				curr.Left = newNode
				newNode.Parent = curr
				break
			}
			curr = curr.Left
		} else {

			if curr.Right == nil {
				curr.Right = newNode
				newNode.Parent = curr
				break
			}
			curr = curr.Right
		}
	}

	return root
}
