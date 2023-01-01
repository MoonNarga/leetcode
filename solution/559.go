package solution

type Node struct {
	Val      int
	Children []*Node
}

func dfs_559(root *Node, depth int) int {
	mdepth := depth
	for _, v := range root.Children {
		mdepth = Max(mdepth, dfs_559(v, depth+1))
	}
	return mdepth
}

func maxDepth_559(root *Node) int {
	if root == nil {
		return 0
	}
	return dfs_559(root, 1)
}
