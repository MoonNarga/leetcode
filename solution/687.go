package solution

func dfs(root *TreeNode, maxLen *int) int {
	if root.Left == nil && root.Right == nil {
		return 1
	}
	currentLen := 1
	longer := 0
	if root.Left != nil {
		longLeft := dfs(root.Left, maxLen)
		if root.Left.Val == root.Val {
			currentLen += longLeft
			longer = longLeft
		}
	}
	if root.Right != nil {
		longRight := dfs(root.Right, maxLen)
		if root.Right.Val == root.Val {
			currentLen += longRight
			if longRight > longer {
				longer = longRight
			}
		}
	}
	if currentLen > *maxLen {
		*maxLen = currentLen
	}
	return 1 + longer
}

func longestUnivaluePath(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxLen := 1
	dfs(root, &maxLen)
	return maxLen - 1
}
