package solution

type Node struct {
	Val      int
	Children []*Node
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func dfs(root *Node, depth int) int {
	mdepth := depth
	for _, v := range root.Children {
		mdepth = max(mdepth, dfs(v, depth+1))
	}
	return mdepth
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	return dfs(root, 1)
}
