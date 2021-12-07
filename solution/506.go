package solution

import (
	"sort"
	"strconv"
)

func findRelativeRanks(score []int) []string {
	co := 1
	rank := make(map[int]string)
	s := make([]int, len(score))
	copy(s, score)
	sort.Ints(s)
	for i := len(s) - 1; i >= 0; i-- {
		switch co {
		case 1:
			{
				rank[s[i]] = "Gold Medal"
			}
		case 2:
			{
				rank[s[i]] = "Silver Medal"
			}
		case 3:
			{
				rank[s[i]] = "Bronze Medal"
			}
		default:
			{
				rank[s[i]] = strconv.Itoa(co)
			}
		}
		co++
	}
	res := make([]string, len(s))
	for i := 0; i < len(s); i++ {
		res[i] = rank[score[i]]
	}
	return res
}
