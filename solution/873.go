package solution

func lenLongestFibSubseq(arr []int) int {
	dp := make([][]int, len(arr))
	m := make(map[int]int, len(arr))
	for i := 0; i < len(arr); i++ {
		dp[i] = make([]int, len(arr))
		m[arr[i]] = i
	}
	maxLen := 0
	for j := 1; j < len(arr); j++ {
		for k := j + 1; k < len(arr); k++ {
			if arr[j]*2 <= arr[k] {
				continue
			}
			if i, ok := m[arr[k]-arr[j]]; ok {
				if dp[i][j] > 2 {
					dp[j][k] = dp[i][j] + 1
				} else {
					dp[j][k] = 3
				}
				if maxLen < dp[j][k] {
					maxLen = dp[j][k]
				}
			}
		}
	}
	return maxLen
}