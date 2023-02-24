package solution

func LongestWPI(hours []int) int {
	indexs := map[int]int{}
	sum, res := 0, 0
	for i, v := range hours {
		if v > 8 {
			sum += 1
		} else {
			sum -= 1
		}
		if sum > 0 {
			res = i + 1
		} else {
			if value, ok := indexs[sum-1]; ok {
				res = Max(res, i-value)
			}
		}
		if _, ok := indexs[sum]; !ok {
			indexs[sum] = i
		}
	}
	return res
}
