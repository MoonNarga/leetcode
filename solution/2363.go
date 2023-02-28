package solution

import "sort"

func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
	m := map[int]int{}
	for _, v := range items2 {
		m[v[0]] = v[1]
	}
	for _, v := range items1 {
		va, ok := m[v[0]]
		if ok {
			v[1] += va
			delete(m, v[0])
		}
	}
	for i, v := range m {
		items1 = append(items1, []int{i, v})
	}
	sort.Slice(items1, func(i, j int) bool {
		return items1[i][0] < items1[j][0]
	})
	return items1
}
