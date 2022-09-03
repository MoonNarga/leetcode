package solution

import "sort"

func findLongestChain(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i][1] < pairs[j][1] {
			return true
		}
		if pairs[i][1] == pairs[j][1] {
			return pairs[i][0] < pairs[j][0]
		}
		return false
	})
	cnt := 0
	end := pairs[0][0] - 1
	for _, v := range pairs {
		if v[0] > end {
			cnt++
			end = v[1]
		}
	}
	return cnt
}
