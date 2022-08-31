package solution

func validateStackSequences(pushed []int, popped []int) bool {
	p := 0
	stack := make([]int, 0)
	for _, v := range pushed {
		stack = append(stack, v)
		for len(stack) > 0 && stack[len(stack)-1] == popped[p] {
			p++
			stack = stack[:len(stack)-1]
		}
	}
	return p == len(popped)
}
