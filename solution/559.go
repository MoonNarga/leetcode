package solution

type Node struct {
	Val      int
	Children []*Node
}

func dfs(root *Node, depth int) int {
	if root == nil {
		return depth
	}
}

func maxDepth(root *Node) int {

}
