package solution

func containOne(root *TreeNode) bool {
	if root == nil {
		return false
	}
	c := 0
	if containOne(root.Left) {
		c++
	} else {
		root.Left = nil
	}
	if containOne(root.Right) {
		c++
	} else {
		root.Right = nil
	}
	if c+root.Val > 0 {
		return true
	}
	return false
}

//using containOne(root) instead of !containOne(root) would increase the memory usage
func pruneTree(root *TreeNode) *TreeNode {
	if !containOne(root) {
		root = nil
	}
	return root
}
