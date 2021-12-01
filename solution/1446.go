package solution

func maxPower(s string) int {
	c := ' '
	maxL := 1
	tL := 1
	for _, v := range s {
		if v == c {
			tL++
			if maxL < tL {
				maxL = tL
			}
		} else {
			tL = 1
			c = v
		}
	}
	return maxL
}
