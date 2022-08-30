package solution

func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	newNode := &TreeNode{Val: val, Left: nil, Right: nil}
	if val > root.Val {
		newNode.Left = root
		return newNode
	}
	parentPoint := root
	for {
		if parentPoint.Right == nil {
			parentPoint.Right = newNode
			break
		} else if val > parentPoint.Right.Val {
			newNode.Left = parentPoint.Right
			parentPoint.Right = newNode
			break
		}
		parentPoint = parentPoint.Right
	}
	return root
}
