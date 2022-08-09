package solution

func minStartValue(nums []int) int {
	currentSum := 0
	minSum := nums[0]
	for _, v := range nums {
		currentSum += v
		if currentSum < minSum {
			minSum = currentSum
		}
	}
	if minSum >= 0 {
		return 1
	}
	return -minSum + 1
}
