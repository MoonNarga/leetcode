package solution

func groupThePeople(groupSizes []int) [][]int {
	n := len(groupSizes)
	res := make([][]int, 0)
	groupOrder := make(map[int]*[]int, 0)
	lenRes := 0
	for i := 0; i < n; i++ {
		if groupSizes[i] == 1 {
			res = append(res, []int{i})
			lenRes++
			continue
		}
		val, ok := groupOrder[groupSizes[i]]
		if ok {
			(*val)[(*val)[0]+1] = i
			(*val)[0]++
			if (*val)[0] == groupSizes[i] {
				res = append(res, make([]int, groupSizes[i]))
				lenRes++
				for j := 0; j < groupSizes[i]; j++ {
					res[lenRes-1][j] = (*val)[j+1]
				}
				(*val)[0] = 0
			}
		} else {
			nGQ := make([]int, groupSizes[i]+1)
			nGQ[0], nGQ[1] = 1, i
			groupOrder[groupSizes[i]] = &nGQ
		}

	}
	return res
}
