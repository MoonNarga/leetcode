package solution

func maxValue(grid [][]int) int {
	dp := make([]int, len(grid[0])+1)
	for i := 0; i < len(grid); i++ {
		for j := 0; j <= len(grid[0]); j++ {
			dp[j] = Max(dp[j], dp[j-1]) + grid[i][j-1]
		}
	}
	return dp[len(grid[0])]
}
