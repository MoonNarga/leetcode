package solution

func constructMaximumBinaryTree(nums []int) *TreeNode {
	stack := make([]*TreeNode, 0)
	for _, v := range nums {
		node := &TreeNode{Val: v, Left: nil, Right: nil}
		for len(stack) != 0 && stack[len(stack)-1].Val < node.Val {
			if node.Left != nil {
				stack[len(stack)-1].Right = node.Left
			}
			node.Left = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, node)
	}
	for i := 0; i < len(stack)-1; i++ {
		stack[i].Right = stack[i+1]
	}
	return stack[0]
}
