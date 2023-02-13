package solution

type col struct {
	h, ind int
}

func trap(height []int) int {
	stack := make([]col, 0)
	i := 0
	for ; i < len(height); i++ {
		if height[i] != 0 {
			break
		}
	}
	sum := 0
	for ; i < len(height); i++ {
		for len(stack) != 0 && height[i] >= stack[len(stack)-1].h {
			if len(stack) == 1 || stack[len(stack)-1].h == height[i] {
				stack = stack[:len(stack)-1]
				break
			}
			sum += (Min(height[i], stack[len(stack)-2].h) - stack[len(stack)-1].h) * (i - stack[len(stack)-2].ind - 1)
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, col{h: height[i], ind: i})
	}
	return sum
}
