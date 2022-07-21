package solution

func containOne(root *TreeNode) bool {
	if root.Left == nil && root.Right == nil {
		if root.Val == 0 {
			return false
		}
		return true
	}
	leftContain, rightContain := false, false
	if root.Left != nil {
		leftContain = containOne(root.Left)
		if leftContain == false {
			root.Left = nil
		}
	}
	if root.Right != nil {
		rightContain = containOne(root.Right)
		if rightContain == false {
			root.Right = nil
		}
	}
	if leftContain == false && rightContain == false && root.Val == 0 {
		return false
	}
	return true
}

func pruneTree(root *TreeNode) *TreeNode {
	if !containOne(root) {
		root = nil
	}
	return root
}
