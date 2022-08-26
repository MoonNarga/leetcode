package solution

func maxProduct_1464(nums []int) int {
	max, subMax := 0, 0
	for _, v := range nums {
		if v > max {
			max, subMax = v, max
		} else if v > subMax {
			subMax = v
		}
	}
	return (max - 1) * (subMax - 1)
}
