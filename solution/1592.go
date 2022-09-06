package solution

import "regexp"

func reorderSpaces(text string) string {
	res := ""
	strs := regexp.MustCompile(`\s+`).Split(text, -1)
	if strs[0] == "" {
		strs = strs[1:]
	}
	if strs[len(strs)-1] == "" {
		strs = strs[:len(strs)-1]
	}
	cnt := 0
	for _, v := range text {
		if v == ' ' {
			cnt++
		}
	}
	if len(strs) == 1 {
		res += strs[0]
		for i := 0; i < cnt; i++ {
			res += " "
		}
		return res
	}
	each := cnt / (len(strs) - 1)
	sp := ""
	for i := 0; i < each; i++ {
		sp += " "
	}
	for i := 0; i < len(strs)-1; i++ {
		res += strs[i] + sp
	}
	res += strs[len(strs)-1]
	for i := 0; i < (cnt - each*(len(strs)-1)); i++ {
		res += " "
	}
	return res
}
