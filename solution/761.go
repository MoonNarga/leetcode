package solution

import "sort"

func makeLargestSpecial(s string) string {
	if s == "" {
		return ""
	}
	subStrings := make([]string, 0)
	cnt := 0
	str := []rune(s)
	start := 0
	for i, v := range str {
		if v == '1' {
			cnt++
		} else {
			cnt--
		}
		if cnt == 0 {
			subStrings = append(subStrings, "1"+makeLargestSpecial(string(str[start+1:i]))+"0")
			start = i + 1
		}
	}
	sort.Slice(subStrings, func(i, j int) bool { return subStrings[i] > subStrings[j] })
	res := ""
	for _, v := range subStrings {
		res += v
	}
	return res
}
