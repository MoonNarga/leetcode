package solution

func maxChunksToSorted(arr []int) int {
	stack := make([]int, len(arr))
	top := 0
	for _, v := range arr {
		if top == 0 {
			stack[top] = v
			top++
		} else {
			if v >= stack[top-1] {
				stack[top] = v
				top++
			} else {
				max := stack[top-1]
				for top > 0 && v < stack[top-1] {
					top--
				}
				stack[top] = max
			}
		}
	}
	return top
}
