package solution

func uniqueLetterString(s string) int {
	str := []rune(s)
	indices := map[rune][]int{}
	for i, v := range str {
		indices[v] = append(indices[v], i)
	}
	res := 0
	for _, v := range indices {
		v = append(append([]int{-1}, v...), len(s))
		for i := 1; i < len(v)-1; i++ {
			res += (v[i] - v[i-1]) * (v[i+1] - v[i])
		}
	}
	return res
}
