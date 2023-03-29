package solution

func countVowelStrings(n int) int {
	dp := []int{1, 1, 1, 1, 1}
	for i := 1; i < n; i++ {
		for j := 0; j < 4; j++ {
			sum := 0
			for k := j; k < 5; k++ {
				sum += dp[k]
			}
			dp[j] = sum
		}
	}
	res := 0
	for i := range dp {
		res += dp[i]
	}
	return res
}
