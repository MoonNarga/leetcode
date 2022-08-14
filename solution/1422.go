package solution

func maxScore(s string) int {
	str := []byte(s)
	c0, c1 := 0, 0
	for _, v := range str {
		if v == '1' {
			c1++
		}
	}
	score := 0
	for _, v := range str[:len(str)-1] {
		if v == '0' {
			c0++
		} else {
			c1--
		}
		if c0+c1 > score {
			score = c0 + c1
		}
	}
	return score
}
