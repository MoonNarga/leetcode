package solution

func findLHS(nums []int) int {
	numSet := make(map[int]int)
	length := 0
	for _, v := range nums {
		if _, ok := numSet[v]; ok {
			numSet[v] += 1
		} else {
			numSet[v] = 1
		}
	}
	for i, v := range numSet {
		if val, ok := numSet[i-1]; ok {
			if val+v > length {
				length = val + v
			}
		}
	}
	return length
}
