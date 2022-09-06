package solution

import "fmt"

func dfs_652(root *TreeNode, record *map[string]int, res *[]*TreeNode) string {
	if root == nil {
		return ""
	}
	pre := fmt.Sprintf("%d(%s)(%s)", root.Val, dfs_652(root.Left, record, res), dfs_652(root.Right, record, res))
	if (*record)[pre] == 1 {
		*res = append(*res, root)
	}
	(*record)[pre]++
	return pre
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	res := make([]*TreeNode, 0)
	record := make(map[string]int, 0)
	dfs_652(root, &record, &res)
	return res
}
