package solution

import "strings"

func isPrefixOfWord(sentence string, searchWord string) int {
	words := strings.Split(sentence, " ")
	for i, v := range words {
		if strings.HasPrefix(v, searchWord) {
			return i + 1
		}
	}
	return -1
}
