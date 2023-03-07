package solution

func MinimumDeletions(s string) int {
	dp, cntB := 0, 0
	for _, v := range s {
		switch v {
		case 'a':
			dp = Min(dp+1, cntB)
		case 'b':
			cntB++
		}
	}
	return dp
}
