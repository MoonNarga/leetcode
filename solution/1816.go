package solution

import (
	"bufio"
	"io"
	"strings"
)

func truncateSentence(s string, k int) string {
	var res string
	reader := bufio.NewReader(strings.NewReader(s))
	co := 0
	for {
		word, err := reader.ReadString(' ')
		if co >= k {
			break
		}
		if err == nil || err == io.EOF {
			res += word
			co++
		}
	}
	return strings.TrimSpace(res)
}
