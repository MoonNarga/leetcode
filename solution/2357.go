package solution

func minimumOperations(nums []int) int {
	m := map[int]bool{}
	for _, v := range nums {
		m[v] = true
	}
	if _, ok := m[0]; ok {
		return len(m) - 1
	}
	return len(m)
}
