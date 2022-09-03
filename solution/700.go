package solution

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	sL := searchBST(root.Left, val)
	if sL != nil {
		return sL
	}
	sR := searchBST(root.Right, val)
	if sR != nil {
		return sR
	}
	return nil
}
