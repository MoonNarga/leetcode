package solution

func constructArray(n int, k int) []int {
	res := make([]int, 0)
	for i, j := 1, 1+k; i < j; i, j = i+1, j-1 {
		res = append(res, i, j)
	}
	if res[len(res)-1]-res[len(res)-2] != 1 {
		res = append(res, res[len(res)-1]-1)
	}
	for i := 2 + k; i <= n; i++ {
		res = append(res, i)
	}
	return res
}
