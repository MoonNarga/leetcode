package solution

func stoneGameVIII(stones []int) int {
	pre := make([]int, len(stones))
	dp := make([]int, len(stones))
	pre[0] = stones[0]
	for i := 1; i < len(stones); i++ {
		pre[i] = stones[i] + pre[i-1]
	}
	dp[len(stones)-1] = pre[len(stones)-1]
	for i := len(stones) - 2; i > 0; i-- {
		dp[i] = maxInt(dp[i+1], pre[i]-dp[i+1])
	}
	return dp[1]
}
