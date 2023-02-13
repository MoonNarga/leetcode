package solution

import "strings"

func getPath(src, dst rune) string {
	if src == dst {
		return "!"
	}
	if src == 'z' {
		return "U" + getPath('u', dst)
	}
	if dst == 'z' {
		temp := getPath(src, 'u')
		return temp[:len(temp)-1] + "D!"
	}
	res := ""
	row := (int(dst) - int('a')) / 5 - (int(src) - int('a')) / 5
	if row < 0 {
		res += strings.Repeat("U", -row)
	} else {
		res += strings.Repeat("D", row)
	}
	col := (int(dst) - int('a')) % 5 - (int(src) - int('a')) % 5
	if col < 0 {
		res += strings.Repeat("L", -col)
	} else {
		res += strings.Repeat("R", col)
	}
	res += "!"
	return res
}

func alphabetBoardPath(target string) string {
	res := ""
	pre := 'a'
	for _, v := range target {
		res += getPath(pre, v)
		pre = v
	}
	return res
}
