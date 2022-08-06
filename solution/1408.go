package solution

import "strings"

func stringMatching(words []string) []string {
	length := len(words)
	res := make([]string, 0)
	flag := make([]bool, length)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			lenA, lenB := len(words[i]), len(words[j])
			if lenA > lenB {
				if flag[j] == false && strings.Contains(words[i], words[j]) {
					res = append(res, words[j])
					flag[j] = true
				}
			} else {
				if flag[i] == false && strings.Contains(words[j], words[i]) {
					res = append(res, words[i])
					flag[i] = true
				}
			}
		}
	}
	return res
}
