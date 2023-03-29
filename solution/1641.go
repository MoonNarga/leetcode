package solution

func countVowelStrings(n int) int {
	dp := []int{1, 1, 1, 1, 1}
	for i := 1; i < n; i++ {
		for j := 0; j < 4; j++ {
			for k := j + 1; k < 5; k++ {
				dp[j] += dp[k]
			}
		}
	}
	return dp[0] + dp[1] + dp[2] + dp[3] + dp[4]
}
