package solution

func verifyPostorder(postorder []int) bool {
	stack := []int{}
	root := int(^uint32(0) >> 1)
	for i := len(postorder) - 1; i >= 0; i-- {
		if postorder[i] > root {
			return false
		}
		if len(stack) == 0 || postorder[i] > stack[len(stack)-1] {
			stack = append(stack, postorder[i])
			continue
		}
		for len(stack) > 0 && stack[len(stack)-1] > postorder[i] {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, postorder[i])
	}
	return true
}
