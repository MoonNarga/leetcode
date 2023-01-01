package solution

func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)
	}
	str1, str2 := []rune(word1), []rune(word2)
	for i := 1; i < len(dp); i++ {
		dp[i][0] = i
	}
	for j := 1; j < len(dp[0]); j++ {
		dp[0][j] = j
	}
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = Min(dp[i-1][j-1]+1, Min(dp[i-1][j]+1, dp[i][j-1]+1))
			}
		}
	}
	return dp[len(word1)][len(word2)]
}
