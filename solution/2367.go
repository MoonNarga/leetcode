package solution

func arithmeticTriplets(nums []int, diff int) (res int) {
	m := map[int]int{}
	for _, v := range nums {
		m[v]++
	}
	for i := range m {
		res += m[i] * m[i+diff] * m[i+2*diff]
	}
	return res
}
